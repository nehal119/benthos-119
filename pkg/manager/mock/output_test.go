package mock_test

import (
	"github.com/nehal119/benthos-119/pkg/component/output"
	"github.com/nehal119/benthos-119/pkg/manager/mock"
)

var _ output.Sync = mock.OutputWriter(nil)
