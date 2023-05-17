package splunk

import (
	_ "embed"

	"github.com/nehal119/benthos-119/pkg/bundle"
	"github.com/nehal119/benthos-119/pkg/template"
)

//go:embed template_output.yaml
var outputTemplate []byte

func init() {
	if err := template.RegisterTemplateYAML(bundle.GlobalEnvironment, outputTemplate); err != nil {
		panic(err)
	}
}
