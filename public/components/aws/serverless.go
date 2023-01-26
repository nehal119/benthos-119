package aws

import (
	"context"

	"github.com/nehal119/benthos-119/internal/serverless/lambda"
)

// RunLambda executes Benthos as an AWS Lambda function. Configuration can be
// stored within the environment variable BENTHOS_CONFIG.
func RunLambda(ctx context.Context) {
	lambda.Run()
}
