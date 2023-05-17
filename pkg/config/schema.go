package config

import (
	"github.com/mitchellh/mapstructure"

	"github.com/nehal119/benthos-119/pkg/api"
	tdocs "github.com/nehal119/benthos-119/pkg/cli/test/docs"
	"github.com/nehal119/benthos-119/pkg/component/metrics"
	"github.com/nehal119/benthos-119/pkg/component/tracer"
	"github.com/nehal119/benthos-119/pkg/docs"
	"github.com/nehal119/benthos-119/pkg/log"
	"github.com/nehal119/benthos-119/pkg/manager"
	"github.com/nehal119/benthos-119/pkg/stream"
)

// Type is the Benthos service configuration struct.
type Type struct {
	HTTP                   api.Config `json:"http" yaml:"http"`
	stream.Config          `json:",inline" yaml:",inline"`
	manager.ResourceConfig `json:",inline" yaml:",inline"`
	Logger                 log.Config     `json:"logger" yaml:"logger"`
	Metrics                metrics.Config `json:"metrics" yaml:"metrics"`
	Tracer                 tracer.Config  `json:"tracer" yaml:"tracer"`
	SystemCloseDelay       string         `json:"shutdown_delay" yaml:"shutdown_delay"`
	SystemCloseTimeout     string         `json:"shutdown_timeout" yaml:"shutdown_timeout"`
	Tests                  []any          `json:"tests,omitempty" yaml:"tests,omitempty"`
}

// New returns a new configuration with default values.
func New() Type {
	return Type{
		HTTP:               api.NewConfig(),
		Config:             stream.NewConfig(),
		ResourceConfig:     manager.NewResourceConfig(),
		Logger:             log.NewConfig(),
		Metrics:            metrics.NewConfig(),
		Tracer:             tracer.NewConfig(),
		SystemCloseDelay:   "",
		SystemCloseTimeout: "20s",
		Tests:              nil,
	}
}

// Clone a config, creating a new copy that can be mutated in isolation.
func (t *Type) Clone() (Type, error) {
	var outConf Type
	if err := mapstructure.Decode(t, &outConf); err != nil {
		return Type{}, err
	}
	return outConf, nil
}

var httpField = docs.FieldObject("http", "Configures the service-wide HTTP server.").WithChildren(api.Spec()...)

var observabilityFields = docs.FieldSpecs{
	docs.FieldObject("logger", "Describes how operational logs should be emitted.").WithChildren(log.Spec()...),
	docs.FieldMetrics("metrics", "A mechanism for exporting metrics.").Optional(),
	docs.FieldTracer("tracer", "A mechanism for exporting traces.").Optional(),
	docs.FieldString("shutdown_delay", "A period of time to wait for metrics and traces to be pulled or pushed from the process.").HasDefault("0s"),
	docs.FieldString("shutdown_timeout", "The maximum period of time to wait for a clean shutdown. If this time is exceeded Benthos will forcefully close.").HasDefault("20s"),
}

// Spec returns a docs.FieldSpec for an entire Benthos configuration.
func Spec() docs.FieldSpecs {
	fields := docs.FieldSpecs{httpField}
	fields = append(fields, stream.Spec()...)
	fields = append(fields, manager.Spec()...)
	fields = append(fields, observabilityFields...)
	fields = append(fields, tdocs.ConfigSpec())
	return fields
}

// SpecWithoutStream describes a stream config without the core stream fields.
func SpecWithoutStream() docs.FieldSpecs {
	fields := docs.FieldSpecs{httpField}
	fields = append(fields, manager.Spec()...)
	fields = append(fields, observabilityFields...)
	fields = append(fields, tdocs.ConfigSpec())
	return fields
}
