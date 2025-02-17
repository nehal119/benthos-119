package stream

import (
	"gopkg.in/yaml.v3"

	"github.com/nehal119/benthos-119/pkg/component/buffer"
	"github.com/nehal119/benthos-119/pkg/component/input"
	"github.com/nehal119/benthos-119/pkg/component/output"
	"github.com/nehal119/benthos-119/pkg/docs"
	"github.com/nehal119/benthos-119/pkg/pipeline"
)

//------------------------------------------------------------------------------

// Config is a configuration struct representing all four layers of a Benthos
// stream.
type Config struct {
	Input    input.Config    `json:"input" yaml:"input"`
	Buffer   buffer.Config   `json:"buffer" yaml:"buffer"`
	Pipeline pipeline.Config `json:"pipeline" yaml:"pipeline"`
	Output   output.Config   `json:"output" yaml:"output"`
}

// NewConfig returns a new configuration with default values.
func NewConfig() Config {
	return Config{
		Input:    input.NewConfig(),
		Buffer:   buffer.NewConfig(),
		Pipeline: pipeline.NewConfig(),
		Output:   output.NewConfig(),
	}
}

// Sanitised returns a sanitised copy of the Benthos configuration, meaning
// fields of no consequence (unused inputs, outputs, processors etc) are
// excluded.
func (c Config) Sanitised() (any, error) {
	var node yaml.Node
	if err := node.Encode(c); err != nil {
		return nil, err
	}

	sanitConf := docs.NewSanitiseConfig()
	sanitConf.RemoveTypeField = true
	if err := Spec().SanitiseYAML(&node, sanitConf); err != nil {
		return nil, err
	}

	var g any
	if err := node.Decode(&g); err != nil {
		return nil, err
	}
	return g, nil
}

//------------------------------------------------------------------------------
