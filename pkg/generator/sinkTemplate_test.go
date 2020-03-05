package generator_test

import (
	"io/ioutil"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func validateSink(tmpDir string, t *testing.T) {
	s, err := ioutil.ReadFile(path.Join(tmpDir, "internal", "kafmesh", "details", "enriched_data_postgres_sink.km.go"))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedSink, string(s))
}

var (
	expectedSink = `// Code generated by kafmesh-gen. DO NOT EDIT.

package details

import (
	"context"
	"time"

	"github.com/lovoo/goka"
	"github.com/pkg/errors"

	"github.com/syncromatics/kafmesh/pkg/runner"

	testSerial "test/internal/kafmesh/models/testMesh/testSerial"
)

type EnrichedDataPostgres_Sink interface {
	Flush() error
	Collect(ctx runner.MessageContext, key string, msg *testSerial.DetailsEnriched) error
}

type impl_EnrichedDataPostgres_Sink struct {
	sink EnrichedDataPostgres_Sink
	codec goka.Codec
	group string
	topic string
	maxBufferSize int
	interval time.Duration
}

func (s *impl_EnrichedDataPostgres_Sink) Codec() goka.Codec {
	return s.codec
}

func (s *impl_EnrichedDataPostgres_Sink) Group() string {
	return s.group
}

func (s *impl_EnrichedDataPostgres_Sink) Topic() string {
	return s.topic
}

func (s *impl_EnrichedDataPostgres_Sink) MaxBufferSize() int {
	return s.maxBufferSize
}

func (s *impl_EnrichedDataPostgres_Sink) Interval() time.Duration {
	return s.interval
}

func (s *impl_EnrichedDataPostgres_Sink) Flush() error {
	return s.sink.Flush()
}

func (s *impl_EnrichedDataPostgres_Sink) Collect(ctx runner.MessageContext, key string, msg interface{}) error {
	m, ok := msg.(*testSerial.DetailsEnriched)
	if !ok {
		return errors.Errorf("expecting message of type '*testSerial.DetailsEnriched' got type '%t'", msg)
	}

	return s.sink.Collect(ctx, key, m)
}

func Register_EnrichedDataPostgres_Sink(options runner.ServiceOptions, sink EnrichedDataPostgres_Sink, interval time.Duration, maxBufferSize int) (func(ctx context.Context) func() error, error) {
	brokers := options.Brokers
	protoWrapper := options.ProtoWrapper

	codec, err := protoWrapper.Codec("testMesh.testSerial.detailsEnriched", &testSerial.DetailsEnriched{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create codec")
	}

	d := &impl_EnrichedDataPostgres_Sink{
		sink: sink,
		codec: codec,
		group: "testService.details.enricheddatapostgres-sink",
		topic: "testMesh.testSerial.detailsEnriched",
		maxBufferSize: maxBufferSize,
		interval: interval,
	}

	s := runner.NewSinkRunner(d, brokers)

	return func(ctx context.Context) func() error {
		return s.Run(ctx)
	}, nil
}
`
)
