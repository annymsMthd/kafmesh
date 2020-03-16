package generator

import (
	"io"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"github.com/syncromatics/kafmesh/pkg/models"
)

var (
	sourceTemplate = template.Must(template.New("").Parse(`// Code generated by kafmesh-gen. DO NOT EDIT.

package {{ .Package }}

import (
	"context"

	"github.com/burdiyan/kafkautil"
	"github.com/lovoo/goka"
	"github.com/pkg/errors"

	"github.com/syncromatics/kafmesh/pkg/runner"

	"{{ .Import }}"
)

type {{ .Name }}_Source interface {
	Emit(message {{ .Name }}_Source_Message) error
	EmitBulk(ctx context.Context, messages []{{ .Name }}_Source_Message) error
	Delete(key string) error
}

type {{ .Name }}_Source_impl struct {
	emitter *runner.Emitter
}

type {{ .Name }}_Source_Message struct {
	Key string
	Value *{{ .MessageType }}
}

type impl_{{ .Name }}_Source_Message struct {
	msg {{ .Name }}_Source_Message
}

func (m *impl_{{ .Name }}_Source_Message) Key() string {
	return m.msg.Key
}

func (m *impl_{{ .Name }}_Source_Message) Value() interface{} {
	return m.msg.Value
}

func New_{{ .Name }}_Source(options runner.ServiceOptions) (*{{ .Name }}_Source_impl, error) {
	brokers := options.Brokers
	protoWrapper := options.ProtoWrapper

	codec, err := protoWrapper.Codec("{{ .TopicName }}", &{{ .MessageType }}{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create codec")
	}

	emitter, err := goka.NewEmitter(brokers,
		goka.Stream("{{ .TopicName }}"),
		codec,
		goka.WithEmitterHasher(kafkautil.MurmurHasher))

	if err != nil {
		return nil, errors.Wrap(err, "failed creating source")
	}

	return &{{ .Name }}_Source_impl{
		emitter: runner.NewEmitter(emitter),
	}, nil
}

func (e *{{ .Name }}_Source_impl) Watch(ctx context.Context) func() error {
	return e.emitter.Watch(ctx)
}

func (e *{{ .Name }}_Source_impl) Emit(message {{ .Name }}_Source_Message) error {
	return e.emitter.Emit(message.Key, message.Value)
}

func (e *{{ .Name }}_Source_impl) EmitBulk(ctx context.Context, messages []{{ .Name }}_Source_Message) error {
	b := []runner.EmitMessage{}
	for _, m := range messages {
		b = append(b, &impl_{{ .Name }}_Source_Message{msg: m})
	}
	return e.emitter.EmitBulk(ctx, b)
}

func (e *{{ .Name }}_Source_impl) Delete(key string) error {
	return e.emitter.Emit(key, nil)
}
`))
)

type sourceOptions struct {
	Package     string
	Import      string
	Name        string
	TopicName   string
	MessageType string
}

func generateSource(writer io.Writer, source *sourceOptions) error {
	err := sourceTemplate.Execute(writer, source)
	if err != nil {
		return errors.Wrap(err, "failed to execute source template")
	}
	return nil
}

func buildSourceOptions(pkg string, mod string, modelsPath string, service *models.Service, source models.Source) (*sourceOptions, error) {
	options := &sourceOptions{
		Package: pkg,
	}

	var name strings.Builder
	nameFrags := strings.Split(source.Message, ".")
	for _, f := range nameFrags[1:] {
		name.WriteString(strcase.ToCamel(f))
	}

	options.TopicName = source.ToTopicName(service)
	options.Name = name.String()
	options.Import = source.ToPackage(service)
	options.MessageType = nameFrags[len(nameFrags)-2] + "." + strcase.ToCamel(nameFrags[len(nameFrags)-1])

	return options, nil
}