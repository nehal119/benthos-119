package mock_test

import (
	"github.com/nehal119/benthos-119/pkg/bundle"
	"github.com/nehal119/benthos-119/pkg/manager/mock"
)

var _ bundle.NewManagement = &mock.Manager{}
