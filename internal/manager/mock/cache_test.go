package mock_test

import (
	"github.com/nehal119/benthos-119/internal/component/cache"
	"github.com/nehal119/benthos-119/internal/manager/mock"
)

var _ cache.V1 = &mock.Cache{}
