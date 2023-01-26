package mock_test

import (
	"github.com/nehal119/benthos-119/internal/component/ratelimit"
	"github.com/nehal119/benthos-119/internal/manager/mock"
)

var _ ratelimit.V1 = mock.RateLimit(nil)
