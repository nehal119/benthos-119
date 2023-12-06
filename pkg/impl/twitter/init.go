package twitter

import (
	_ "embed"
	// bloblang functions are registered in init functions under this package
	// so ensure they are loaded first
	_ "github.com/nehal119/benthos-119/pkg/impl/pure"

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
