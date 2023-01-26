package pure

import (
	"go.opentelemetry.io/otel/trace"

	"github.com/nehal119/benthos-119/pkg/bundle"
	"github.com/nehal119/benthos-119/pkg/component/tracer"
	"github.com/nehal119/benthos-119/pkg/docs"
)

func init() {
	_ = bundle.AllTracers.Add(func(c tracer.Config, nm bundle.NewManagement) (trace.TracerProvider, error) {
		return trace.NewNoopTracerProvider(), nil
	}, docs.ComponentSpec{
		Name:    "none",
		Type:    docs.TypeTracer,
		Status:  docs.StatusStable,
		Summary: `Do not send tracing events anywhere.`,
		Config:  docs.FieldObject("", ""),
	})
}
