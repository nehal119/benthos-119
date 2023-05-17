package twitter

import (
	_ "embed"

	"github.com/nehal119/benthos-119/pkg/bundle"
	"github.com/nehal119/benthos-119/pkg/template"
)

//go:embed template_search_input.yaml
var searchInputTemplate []byte

func init() {
	if err := template.RegisterTemplateYAML(bundle.GlobalEnvironment, searchInputTemplate); err != nil {
		panic(err)
	}
}
