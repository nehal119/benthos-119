package pure

import (
	"context"
	"fmt"

	"github.com/nehal119/benthos-119/pkg/bloblang/field"
	"github.com/nehal119/benthos-119/pkg/bundle"
	"github.com/nehal119/benthos-119/pkg/component/processor"
	"github.com/nehal119/benthos-119/pkg/docs"
	"github.com/nehal119/benthos-119/pkg/log"
	"github.com/nehal119/benthos-119/pkg/message"
)

func init() {
	err := bundle.AllProcessors.Add(func(conf processor.Config, mgr bundle.NewManagement) (processor.V1, error) {
		p, err := newGroupByValue(conf.GroupByValue, mgr)
		if err != nil {
			return nil, err
		}
		return processor.NewAutoObservedBatchedProcessor("group_by_value", p, mgr), nil
	}, docs.ComponentSpec{
		Name: "group_by_value",
		Categories: []string{
			"Composition",
		},
		Summary: `Splits a batch of messages into N batches, where each resulting batch contains a group of messages determined by a [function interpolated string](/docs/configuration/interpolation#bloblang-queries) evaluated per message.`,
		Description: `
This allows you to group messages using arbitrary fields within their content or metadata, process them individually, and send them to unique locations as per their group.

The functionality of this processor depends on being applied across messages that are batched. You can find out more about batching [in this doc](/docs/configuration/batching).`,
		Footnotes: `
## Examples

If we were consuming Kafka messages and needed to group them by their key,
archive the groups, and send them to S3 with the key as part of the path we
could achieve that with the following:

` + "```yaml" + `
pipeline:
  processors:
    - group_by_value:
        value: ${! meta("kafka_key") }
    - archive:
        format: tar
    - compress:
        algorithm: gzip
output:
  aws_s3:
    bucket: TODO
    path: docs/${! meta("kafka_key") }/${! count("files") }-${! timestamp_unix_nano() }.tar.gz
` + "```" + ``,
		Config: docs.FieldComponent().WithChildren(
			docs.FieldString(
				"value", "The interpolated string to group based on.",
				"${! meta(\"kafka_key\") }", "${! json(\"foo.bar\") }-${! meta(\"baz\") }",
			).IsInterpolated().HasDefault(""),
		),
	})
	if err != nil {
		panic(err)
	}
}

type groupByValueProc struct {
	log   log.Modular
	value *field.Expression
}

func newGroupByValue(conf processor.GroupByValueConfig, mgr bundle.NewManagement) (processor.AutoObservedBatched, error) {
	value, err := mgr.BloblEnvironment().NewField(conf.Value)
	if err != nil {
		return nil, fmt.Errorf("failed to parse value expression: %v", err)
	}
	return &groupByValueProc{
		log:   mgr.Logger(),
		value: value,
	}, nil
}

func (g *groupByValueProc) ProcessBatch(ctx *processor.BatchProcContext, batch message.Batch) ([]message.Batch, error) {
	if batch.Len() == 0 {
		return nil, nil
	}

	groupKeys := []string{}
	groupMap := map[string]message.Batch{}

	_ = batch.Iter(func(i int, p *message.Part) error {
		v, err := g.value.String(i, batch)
		if err != nil {
			g.log.Errorf("Group value interpolation error: %v", err)
			err = fmt.Errorf("group value interpolation error: %w", err)
			ctx.OnError(err, i, p)
		}

		ctx.Span(i).LogKV("event", "grouped", "type", v)
		ctx.Span(i).SetTag("group", v)
		if group, exists := groupMap[v]; exists {
			groupMap[v] = append(group, p)
		} else {
			g.log.Tracef("New group formed: %v\n", v)
			groupKeys = append(groupKeys, v)
			groupMap[v] = message.Batch{p}
		}
		return nil
	})

	msgs := []message.Batch{}
	for _, key := range groupKeys {
		msgs = append(msgs, groupMap[key])
	}
	if len(msgs) == 0 {
		return nil, nil
	}
	return msgs, nil
}

func (g *groupByValueProc) Close(context.Context) error {
	return nil
}
