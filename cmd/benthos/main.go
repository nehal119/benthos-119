package main

import (
	"context"

	"github.com/nehal119/benthos-119/pkg/cli"

	// Import all plugins defined within the repo.
	_ "github.com/nehal119/benthos-119/public/components/all"
)

func main() {
	cli.Run(context.Background())
}
