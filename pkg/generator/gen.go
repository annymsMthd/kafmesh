package generator

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pkg/errors"
	"github.com/syncromatics/kafmesh/pkg/models"
	"github.com/yargevad/filepathx"
)

// Options for the kafmesh generator
type Options struct {
	Service         *models.Service
	Components      []*models.Component
	RootPath        string
	DefinitionsPath string
}

// Generate generates the kafmesh files
func Generate(options Options) error {
	outputPath := path.Join(options.RootPath, options.Service.Output.Path)

	err := os.MkdirAll(outputPath, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "failed to create output path")
	}

	includes := []string{}
	files := []file{}
	for _, p := range options.Service.Messages.Protobuf {
		protoPath := p
		if runtime.GOOS == "windows" {
			protoPath = strings.ReplaceAll(protoPath, "/", "\\")
		}
		protoPath = path.Join(options.DefinitionsPath, protoPath)
		if runtime.GOOS == "windows" {
			protoPath = strings.ReplaceAll(protoPath, "/", "\\")
		}
		protoPath, err = filepath.Abs(protoPath)
		if err != nil {
			return errors.Wrapf(err, "failed to get absolute path from protopath '%s'", protoPath)
		}

		includes = append(includes, protoPath)

		fs, err := filepathx.Glob(path.Join(protoPath, "**/*.proto"))
		if err != nil {
			return errors.Wrap(err, "failed to glob files")
		}

		if len(fs) == 0 {
			return errors.Errorf("no proto files found in '%s'", protoPath)
		}

		for _, f := range fs {
			if runtime.GOOS == "windows" {
				f = strings.ReplaceAll(f, "/", "\\")
			}
			files = append(files, file{
				root: protoPath,
				path: f,
			})
		}
	}

	modelsPath := path.Join(options.Service.Output.Path, "models")
	err = os.MkdirAll(path.Join(options.RootPath, modelsPath), os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "failed to create output models path")
	}

	protocOptions := protoOptions{
		Files:    files,
		Includes: includes,
		Output:   path.Join(options.RootPath, modelsPath),
		Module:   options.Service.Output.Module,
		Path:     options.Service.Output.Path,
	}

	err = Protoc(protocOptions)
	if err != nil {
		return errors.Wrapf(err, "failed to run protoc")
	}

	file, err := os.Create(path.Join(outputPath, "service.km.go"))
	if err != nil {
		return errors.Wrapf(err, "failed to open service file")
	}
	defer file.Close()

	for _, c := range options.Components {
		err = processComponent(options.RootPath, outputPath, options.Service.Output.Module, modelsPath, options.Service, c)
		if err != nil {
			return errors.Wrapf(err, "failed to process component")
		}
	}

	sOptions, err := buildServiceOptions(options.Service, options.Components, options.Service.Output.Module)
	if err != nil {
		return errors.Wrapf(err, "failed to build options")
	}

	err = generateService(file, sOptions)
	if err != nil {
		return errors.Wrapf(err, "failed to generate package")
	}

	tOptions, err := buildTopicOption(options.Service, options.Components)
	if err != nil {
		return errors.Wrap(err, "failed to build topic options")
	}

	file, err = os.Create(path.Join(outputPath, "topics.km.go"))
	if err != nil {
		return errors.Wrapf(err, "failed to open topics file")
	}
	defer file.Close()

	err = generateTopics(file, tOptions)
	if err != nil {
		return errors.Wrapf(err, "failed to generate topics")
	}

	err = generateMocks(outputPath)
	if err != nil {
		return errors.Wrapf(err, "failed to generate mocks")
	}

	return nil
}

func processComponent(rootPath string, outputPath string, mod string, modelsPath string, service *models.Service, component *models.Component) error {
	mPath := "/" + strings.TrimPrefix(modelsPath, rootPath)
	componentPath := path.Join(outputPath, component.Name)
	_, err := os.Stat(componentPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(componentPath, os.ModePerm)
		if err != nil {
			return errors.Wrap(err, "failed to create component path")
		}
	}

	for _, p := range component.Processors {
		fileName := p.Name
		fileName = fmt.Sprintf("%s_processor.km.go", fileName)
		file, err := os.Create(path.Join(componentPath, fileName))
		if err != nil {
			return errors.Wrapf(err, "failed to open service file")
		}
		defer file.Close()

		co, err := buildProcessorOptions(component.Name, mod, mPath, service, component, p)
		if err != nil {
			return errors.Wrap(err, "failed to build processor options")
		}

		err = generateProcessor(file, co)
		if err != nil {
			return errors.Wrap(err, "failed to generate processor")
		}
	}

	for _, e := range component.Emitters {
		fileName := strings.ReplaceAll(e.Message, ".", "_")
		fileName = fmt.Sprintf("%s_emitter.km.go", fileName)
		file, err := os.Create(path.Join(componentPath, fileName))
		if err != nil {
			return errors.Wrapf(err, "failed to open service file")
		}
		defer file.Close()

		co, err := buildEmitterOptions(component.Name, mod, mPath, service, e)
		if err != nil {
			return errors.Wrap(err, "failed to build emitter options")
		}

		err = generateEmitter(file, co)
		if err != nil {
			return errors.Wrap(err, "failed to generate emitter")
		}
	}

	for _, s := range component.Sinks {
		fileName := strings.ReplaceAll(s.Name, " ", "_")
		fileName = fmt.Sprintf("%s_sink.km.go", fileName)
		fileName = strings.ToLower(fileName)
		file, err := os.Create(path.Join(componentPath, fileName))
		if err != nil {
			return errors.Wrapf(err, "failed to open service file")
		}
		defer file.Close()

		co, err := buildSinkOptions(component.Name, mod, mPath, s, service, component)
		if err != nil {
			return errors.Wrap(err, "failed to build sink options")
		}

		err = generateSink(file, co)
		if err != nil {
			return errors.Wrap(err, "failed to generate sink")
		}
	}

	for _, v := range component.Views {
		fileName := strings.ReplaceAll(v.Message, ".", "_")
		fileName = fmt.Sprintf("%s_view.km.go", fileName)
		file, err := os.Create(path.Join(componentPath, fileName))
		if err != nil {
			return errors.Wrapf(err, "failed to open service file")
		}
		defer file.Close()

		co, err := buildViewOptions(component.Name, mod, mPath, service, v)
		if err != nil {
			return errors.Wrap(err, "failed to build view options")
		}

		err = generateView(file, co)
		if err != nil {
			return errors.Wrap(err, "failed to generate view")
		}
	}

	for _, s := range component.Synchronizers {
		fileName := strings.ReplaceAll(s.Message, ".", "_")
		fileName = fmt.Sprintf("%s_synchronizer.km.go", fileName)
		file, err := os.Create(path.Join(componentPath, fileName))
		if err != nil {
			return errors.Wrapf(err, "failed to open service file")
		}
		defer file.Close()

		co, err := buildSynchronizerOptions(component.Name, mod, mPath, service, s)
		if err != nil {
			return errors.Wrap(err, "failed to build synchronizer options")
		}

		err = generateSynchronizer(file, co)
		if err != nil {
			return errors.Wrap(err, "failed to generate synchronizer")
		}
	}

	return nil
}
