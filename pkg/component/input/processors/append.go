package processors

import (
	"fmt"
	"strconv"

	"github.com/nehal119/benthos-119/pkg/bundle"
	"github.com/nehal119/benthos-119/pkg/component/input"
	"github.com/nehal119/benthos-119/pkg/component/processor"
	"github.com/nehal119/benthos-119/pkg/pipeline"
)

// AppendFromConfig takes a variant arg of pipeline constructor functions and
// returns a new slice of them where the processors of the provided input
// configuration will also be initialized.
func AppendFromConfig(conf input.Config, mgr bundle.NewManagement, pipelines ...processor.PipelineConstructorFunc) []processor.PipelineConstructorFunc {
	if len(conf.Processors) > 0 {
		pipelines = append([]processor.PipelineConstructorFunc{func() (processor.Pipeline, error) {
			processors := make([]processor.V1, len(conf.Processors))
			for j, procConf := range conf.Processors {
				newMgr := mgr.IntoPath("processors", strconv.Itoa(j))
				var err error
				processors[j], err = newMgr.NewProcessor(procConf)
				if err != nil {
					return nil, fmt.Errorf("failed to create processor '%v': %v", procConf.Type, err)
				}
			}
			return pipeline.NewProcessor(processors...), nil
		}}, pipelines...)
	}
	return pipelines
}

// WrapConstructor provides a way to define an input constructor without
// manually initializing processors of the config.
func WrapConstructor(fn func(input.Config, bundle.NewManagement) (input.Streamed, error)) bundle.InputConstructor {
	return func(c input.Config, nm bundle.NewManagement) (input.Streamed, error) {
		i, err := fn(c, nm)
		if err != nil {
			return nil, err
		}
		pcf := AppendFromConfig(c, nm)
		return input.WrapWithPipelines(i, pcf...)
	}
}
