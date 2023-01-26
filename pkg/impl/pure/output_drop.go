package pure

import (
	"context"

	"github.com/nehal119/benthos-119/pkg/bundle"
	"github.com/nehal119/benthos-119/pkg/component/output"
	"github.com/nehal119/benthos-119/pkg/component/output/processors"
	"github.com/nehal119/benthos-119/pkg/docs"
	"github.com/nehal119/benthos-119/pkg/log"
	"github.com/nehal119/benthos-119/pkg/message"
)

func init() {
	err := bundle.AllOutputs.Add(processors.WrapConstructor(func(c output.Config, nm bundle.NewManagement) (output.Streamed, error) {
		return output.NewAsyncWriter("drop", 1, newDropWriter(c.Drop, nm.Logger()), nm)
	}), docs.ComponentSpec{
		Name:       "drop",
		Summary:    `Drops all messages.`,
		Categories: []string{"Utility"},
		Config:     docs.FieldObject("", "").HasDefault(struct{}{}),
	})
	if err != nil {
		panic(err)
	}
}

type dropWriter struct {
	log log.Modular
}

func newDropWriter(conf output.DropConfig, log log.Modular) *dropWriter {
	return &dropWriter{log: log}
}

func (d *dropWriter) Connect(ctx context.Context) error {
	d.log.Infoln("Dropping messages.")
	return nil
}

func (d *dropWriter) WriteBatch(ctx context.Context, msg message.Batch) error {
	return nil
}

func (d *dropWriter) Close(context.Context) error {
	return nil
}
