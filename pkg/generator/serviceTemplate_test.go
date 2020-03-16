package generator_test

import (
	"io/ioutil"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func validateService(tmpDir string, t *testing.T) {
	s, err := ioutil.ReadFile(path.Join(tmpDir, "internal", "kafmesh", "service.km.go"))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedService, string(s))
}

var (
	expectedService = `// Code generated by kafmesh-gen. DO NOT EDIT.

package kafmesh

import (
	"time"

	"github.com/pkg/errors"
	"github.com/syncromatics/kafmesh/pkg/runner"

	"test/internal/kafmesh/details"
)

func Register_Details_Enricher_Processor(service *runner.Service, processor details.Enricher_Processor) error {
	r, err := details.Register_Enricher_Processor(service.Options(), processor)
	if err != nil {
		return errors.Wrap(err, "failed to register processor")
	}

	err = service.RegisterRunner(r)
	if err != nil {
		return errors.Wrap(err, "failed to register runner with service")
	}

	return nil
}

func New_TestSerialDetails_Source(service *runner.Service) (details.TestSerialDetails_Source, error) {
	e, err := details.New_TestSerialDetails_Source(service.Options())
	if err != nil {
		return nil, err
	}

	err = service.RegisterRunner(e.Watch)
	if err != nil {
		return nil, errors.Wrap(err, "failed to register runner with service")
	}

	return e, nil
}

func New_TestSerialDetailsEnriched_View(service *runner.Service) (details.TestSerialDetailsEnriched_View, error) {
	v, err := details.New_TestSerialDetailsEnriched_View(service.Options())
	if err != nil {
		return nil, err
	}

	err = service.RegisterRunner(v.Watch)
	if err != nil {
		return nil, errors.Wrap(err, "failed to register runner with service")
	}

	return v, nil
}

func Register_EnrichedDataPostgres_Sink(service *runner.Service, sink details.EnrichedDataPostgres_Sink, interval time.Duration, maxBufferSize int) error {
	r, err := details.Register_EnrichedDataPostgres_Sink(service.Options(), sink, interval, maxBufferSize)
	if err != nil {
		return errors.Wrap(err, "failed to register sink")
	}

	err = service.RegisterRunner(r)
	if err != nil {
		return errors.Wrap(err, "failed to register runner with service")
	}

	return nil
}

func Register_Details_TestToDatabase_ViewSource(service *runner.Service, viewSource details.TestToDatabase_ViewSource, updateInterval time.Duration, syncTimeout time.Duration) error {
	r, err := details.Register_TestToDatabase_ViewSource(service.Options(), viewSource, updateInterval, syncTimeout)
	if err != nil {
		return errors.Wrap(err, "failed to register viewSource")
	}

	err = service.RegisterRunner(r)
	if err != nil {
		return errors.Wrap(err, "failed to register runner with service")
	}

	return nil
}

func Register_Details_TestToApi_ViewSink(service *runner.Service, viewSink details.TestToApi_ViewSink, updateInterval time.Duration, syncTimeout time.Duration) error {
	r, err := details.Register_TestToApi_ViewSink(service.Options(), viewSink, updateInterval, syncTimeout)
	if err != nil {
		return errors.Wrap(err, "failed to register viewSink")
	}

	err = service.RegisterRunner(r)
	if err != nil {
		return errors.Wrap(err, "failed to register runner with service")
	}

	return nil
}
`
)
