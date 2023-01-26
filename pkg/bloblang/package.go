package bloblang

import (
	"github.com/nehal119/benthos-119/pkg/bloblang/plugins"
)

func init() {
	if err := plugins.Register(); err != nil {
		panic(err)
	}
}
