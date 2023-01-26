package pure

import (
	"github.com/nehal119/benthos-119/pkg/bundle"
	"github.com/nehal119/benthos-119/pkg/component/metrics"
	"github.com/nehal119/benthos-119/pkg/docs"
)

func init() {
	_ = bundle.AllMetrics.Add(func(metrics.Config, bundle.NewManagement) (metrics.Type, error) {
		return metrics.Noop(), nil
	}, docs.ComponentSpec{
		Name:    "none",
		Type:    docs.TypeMetrics,
		Summary: `Disable metrics entirely.`,
		Config:  docs.FieldObject("", "").HasDefault(struct{}{}),
	})
}
