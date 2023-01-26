package mock_test

import (
	"github.com/nehal119/benthos-119/internal/component/processor"
	"github.com/nehal119/benthos-119/internal/manager/mock"
)

var _ processor.V1 = mock.Processor(nil)
