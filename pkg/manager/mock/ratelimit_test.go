package mock_test

import (
	"github.com/nehal119/benthos-119/pkg/component/ratelimit"
	"github.com/nehal119/benthos-119/pkg/manager/mock"
)

var _ ratelimit.V1 = mock.RateLimit(nil)
