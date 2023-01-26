package mock_test

import (
	"github.com/nehal119/benthos-119/pkg/component/input"
	"github.com/nehal119/benthos-119/pkg/manager/mock"
)

var _ input.Streamed = &mock.Input{}
