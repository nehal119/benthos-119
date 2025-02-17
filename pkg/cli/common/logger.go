package common

import (
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/nehal119/benthos-119/pkg/config"
	"github.com/nehal119/benthos-119/pkg/filepath/ifs"
	"github.com/nehal119/benthos-119/pkg/log"
)

// CreateLogger from a CLI context and a stream config.
func CreateLogger(c *cli.Context, conf config.Type, streamsMode bool) (logger log.Modular, err error) {
	if overrideLogLevel := c.String("log.level"); len(overrideLogLevel) > 0 {
		conf.Logger.LogLevel = strings.ToUpper(overrideLogLevel)
	}

	defaultStream := os.Stdout
	if !streamsMode && conf.Output.Type == "stdout" {
		defaultStream = os.Stderr
	}
	logger, err = log.New(defaultStream, ifs.OS(), conf.Logger)
	return
}
