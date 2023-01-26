package mock_test

import (
	"github.com/nehal119/benthos-119/pkg/component/processor"
	"github.com/nehal119/benthos-119/pkg/manager/mock"
)

var _ processor.V1 = mock.Processor(nil)
