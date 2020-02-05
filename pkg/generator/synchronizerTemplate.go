package generator

import (
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"github.com/syncromatics/kafmesh/pkg/models"
)

var (
	synchronizerTemplate = template.Must(template.New("").Parse(`// Code generated by kafmesh-gen. DO NOT EDIT.

package {{ .Package }}

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/burdiyan/kafkautil"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/storage"
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"golang.org/x/sync/errgroup"

	"github.com/syncromatics/kafmesh/pkg/runner"

	{{ .Import }}
)

type {{ .Name }}_Synchronizer_Message struct {
	key string
	value *{{ .MessageType }}
}

func New_{{ .Name }}_Synchronizer_Message(key string, value *{{ .MessageType }}) *{{ .Name }}_Synchronizer_Message {
	return &{{ .Name }}_Synchronizer_Message{
		key: key,
		value: value,
	}
}

func (m *{{ .Name }}_Synchronizer_Message) Key() string {
	return m.key
}

func (m *{{ .Name }}_Synchronizer_Message) Value() interface{} {
	return m.value
}

type {{ .Name }}_Synchronizer_Context interface {
	Keys() ([]string, error)
	Get(string) (*{{ .MessageType }}, error)
	Emit(*{{ .Name }}_Synchronizer_Message) error
	EmitBulk(context.Context, []*{{ .Name }}_Synchronizer_Message) error
}

type {{ .Name }}_Synchronizer_Context_impl struct {
	view *goka.View
	emitter *runner.Emitter
}

func (c *{{ .Name }}_Synchronizer_Context_impl) Keys() ([]string, error) {
	it, err := c.view.Iterator()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get iterator")
	}

	keys := []string{}
	for it.Next() {
		keys = append(keys, it.Key())
	}

	return keys, nil
}

func (c *{{ .Name }}_Synchronizer_Context_impl) Get(key string) (*{{ .MessageType }}, error) {
	m, err := c.view.Get(key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get value from view")
	}

	msg, ok := m.(*{{ .MessageType }})
	if !ok {
		return nil, errors.Errorf("expecting message of type '*{{ .MessageType }}' got type '%t'", m)
	}

	return msg, nil
}

func (c *{{ .Name }}_Synchronizer_Context_impl) Emit(message *{{ .Name }}_Synchronizer_Message) error {
	return c.emitter.Emit(message.Key(), message.Value())
}

func (c *{{ .Name }}_Synchronizer_Context_impl) EmitBulk(ctx context.Context, messages []*{{ .Name }}_Synchronizer_Message) error {
	b := []runner.EmitMessage{}
	for _, m := range messages {
		b = append(b, m)
	}
	return c.emitter.EmitBulk(ctx, b)
}

type {{ .Name }}_Synchronizer interface {
	Sync({{ .Name }}_Synchronizer_Context) error
}

func Register_{{ .Name }}_Synchronizer(options runner.ServiceOptions, sychronizer {{ .Name }}_Synchronizer, updateInterval time.Duration) (func(context.Context) func() error, error) {
	brokers := options.Brokers
	protoWrapper := options.ProtoWrapper

	codec, err := protoWrapper.Codec("{{ .TopicName }}", &{{ .MessageType }}{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create codec")
	}

	opts := &opt.Options{
		BlockCacheCapacity: opt.MiB * 1,
		WriteBuffer:        opt.MiB * 1,
	}

	path := filepath.Join("/tmp/storage", "synchronizer", "{{ .TopicName }}")

	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create synchronizer db directory")
	}

	builder := storage.BuilderWithOptions(path, opts)

	view, err := goka.NewView(brokers,
		goka.Table("{{ .TopicName }}"),
		codec,
		goka.WithViewStorageBuilder(builder),
		goka.WithViewHasher(kafkautil.MurmurHasher),
	)

	if err != nil {
		return nil, errors.Wrap(err, "failed creating sychronizer view")
	}

	e, err := goka.NewEmitter(brokers,
		goka.Stream("{{ .TopicName }}"),
		codec,
		goka.WithEmitterHasher(kafkautil.MurmurHasher))

	if err != nil {
		return nil, errors.Wrap(err, "failed creating sychronizer emitter")
	}

	emitter := runner.NewEmitter(e)

	c := &{{ .Name }}_Synchronizer_Context_impl{
		view:    view,
		emitter: emitter,
	}

	return func(ctx context.Context) func() error {
		return func() error {
			gctx, cancel := context.WithCancel(ctx)
			grp, gctx := errgroup.WithContext(ctx)

			timer := time.NewTimer(0)
			grp.Go(func() error {
				for {
					select {
					case <-gctx.Done():
						return nil
					case <-timer.C:
						err := sychronizer.Sync(c)
						if err != nil {
							return err
						}
						timer = time.NewTimer(updateInterval)
					}
				}
			})

			grp.Go(emitter.Watch(gctx))
			grp.Go(func() error {
				return view.Run(ctx)
			})

			select {
			case <- ctx.Done():
				cancel()
				return nil
			case <- gctx.Done():
				cancel()
				err := grp.Wait()
				return err
			}
		}
	}, nil
}
`))
)

type synchronizerOptions struct {
	Package     string
	Import      string
	Name        string
	TopicName   string
	MessageType string
}

func generateSynchronizer(writer io.Writer, synchronizer *synchronizerOptions) error {
	err := synchronizerTemplate.Execute(writer, synchronizer)
	if err != nil {
		return errors.Wrap(err, "failed to execute synchronizer template")
	}
	return nil
}

func buildSynchronizerOptions(pkg string, mod string, modelsPath string, synchronizer models.Synchronizer) (*synchronizerOptions, error) {
	options := &synchronizerOptions{
		Package: pkg,
	}

	var name strings.Builder
	nameFrags := strings.Split(synchronizer.Message, ".")
	for _, f := range nameFrags[1:] {
		name.WriteString(strcase.ToCamel(f))
	}

	topic := synchronizer.Message
	if synchronizer.TopicDefinition.Topic != nil {
		topic = *synchronizer.TopicDefinition.Topic
	}

	options.TopicName = topic
	options.Name = name.String()

	var mPkg strings.Builder
	for _, p := range nameFrags[:len(nameFrags)-1] {
		mPkg.WriteString("/")
		mPkg.WriteString(p)
	}

	imp := strings.TrimPrefix(mPkg.String(), "/")
	d := strings.Split(imp, "/")
	options.Import = fmt.Sprintf("%s \"%s%s/%s\"", d[len(d)-1], mod, modelsPath, imp)

	options.MessageType = nameFrags[len(nameFrags)-2] + "." + strcase.ToCamel(nameFrags[len(nameFrags)-1])

	return options, nil
}
