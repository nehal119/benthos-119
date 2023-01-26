package mock_test

import (
	"github.com/nehal119/benthos-119/internal/bundle"
	"github.com/nehal119/benthos-119/internal/manager/mock"
)

var _ bundle.NewManagement = &mock.Manager{}
