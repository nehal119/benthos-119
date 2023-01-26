package mock_test

import (
	"github.com/nehal119/benthos-119/pkg/component/cache"
	"github.com/nehal119/benthos-119/pkg/manager/mock"
)

var _ cache.V1 = &mock.Cache{}
