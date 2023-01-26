package main

import (
	"github.com/nehal119/benthos-119/pkg/serverless/lambda"

	// Import all plugins defined within the repo.
	_ "github.com/nehal119/benthos-119/public/components/all"
)

func main() {
	lambda.Run()
}
