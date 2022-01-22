package {{ .Package }}

import (
	"context"
	"time"

	"github.com/syncromatics/kafmesh/pkg/runner"
)

func ConfigureTopics(ctx context.Context, brokers []string) error {
	topics := []runner.Topic{}

	topic := runner.Topic{}

	maxBytes := 0

	{{- range .Topics }}
	topic = runner.Topic{
		Name:       "{{ .Name }}",
		Partitions: {{ .Partitions}},
		Replicas:   {{ .Replicas }},
		Compact:    {{ .Compact }},
		Retention:  {{ .Retention.Milliseconds }} * time.Millisecond,
		Segment:    {{ .Segment.Milliseconds }} * time.Millisecond,
		Create:     {{ .Create }},
	}

	{{- if not .MaxMessageBytes -}}
	{{- else }}
	maxBytes = {{ .MaxMessageBytes }}
	topic.MaxMessageBytes = &maxBytes
	{{- end }}
	topics = append(topics, topic)
	{{ end }}

	return runner.ConfigureTopics(ctx, brokers, topics)
}
