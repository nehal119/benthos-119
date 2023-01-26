package mock_test

import (
	"github.com/nehal119/benthos-119/internal/component/input"
	"github.com/nehal119/benthos-119/internal/manager/mock"
)

var _ input.Streamed = &mock.Input{}
