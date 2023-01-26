package mock_test

import (
	"github.com/nehal119/benthos-119/internal/component/output"
	"github.com/nehal119/benthos-119/internal/manager/mock"
)

var _ output.Sync = mock.OutputWriter(nil)
