package manager

import (
	"errors"

	"github.com/Jeffail/gabs/v2"

	"github.com/nehal119/benthos-119/pkg/docs"
)

func lintResource(ctx docs.LintContext, line, col int, v any) []docs.Lint {
	if _, ok := v.(map[string]any); !ok {
		return nil
	}
	gObj := gabs.Wrap(v)
	label, _ := gObj.S("label").Data().(string)
	if label == "" {
		return []docs.Lint{
			docs.NewLintError(line, docs.LintBadLabel, errors.New("the label field for resources must be unique and not empty")),
		}
	}
	return nil
}

// Spec returns a field spec for the manager configuration.
func Spec() docs.FieldSpecs {
	return docs.FieldSpecs{
		docs.FieldInput(
			"input_resources", "A list of input resources, each must have a unique label.",
		).Array().LinterFunc(lintResource).HasDefault([]any{}),

		docs.FieldProcessor(
			"processor_resources", "A list of processor resources, each must have a unique label.",
		).Array().LinterFunc(lintResource).HasDefault([]any{}),

		docs.FieldOutput(
			"output_resources", "A list of output resources, each must have a unique label.",
		).Array().LinterFunc(lintResource).HasDefault([]any{}),

		docs.FieldCache(
			"cache_resources", "A list of cache resources, each must have a unique label.",
		).Array().LinterFunc(lintResource).HasDefault([]any{}),

		docs.FieldRateLimit(
			"rate_limit_resources", "A list of rate limit resources, each must have a unique label.",
		).Array().LinterFunc(lintResource).HasDefault([]any{}),
	}
}
