package pure

import (
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"

	"github.com/nehal119/benthos-119/pkg/bundle"
	"github.com/nehal119/benthos-119/pkg/component/tracer"
	"github.com/nehal119/benthos-119/pkg/docs"
)

func init() {
	_ = bundle.AllTracers.Add(func(c tracer.Config, nm bundle.NewManagement) (trace.TracerProvider, error) {
		return noop.NewTracerProvider(), nil
	}, docs.ComponentSpec{
		Name:    "none",
		Type:    docs.TypeTracer,
		Status:  docs.StatusStable,
		Summary: `Do not send tracing events anywhere.`,
		Config:  docs.FieldObject("", ""),
	})
}
