package output

import (
	"github.com/nehal119/benthos-119/pkg/batch/policy/batchconfig"
	sess "github.com/nehal119/benthos-119/pkg/impl/aws/session"
	"github.com/nehal119/benthos-119/pkg/metadata"
	"github.com/nehal119/benthos-119/pkg/old/util/retries"
)

// AmazonSQSConfig contains configuration fields for the output AmazonSQS type.
type AmazonSQSConfig struct {
	SessionConfig          `json:",inline" yaml:",inline"`
	URL                    string                       `json:"url" yaml:"url"`
	MessageGroupID         string                       `json:"message_group_id" yaml:"message_group_id"`
	MessageDeduplicationID string                       `json:"message_deduplication_id" yaml:"message_deduplication_id"`
	Metadata               metadata.ExcludeFilterConfig `json:"metadata" yaml:"metadata"`
	MaxInFlight            int                          `json:"max_in_flight" yaml:"max_in_flight"`
	retries.Config         `json:",inline" yaml:",inline"`
	Batching               batchconfig.Config `json:"batching" yaml:"batching"`
}

// NewAmazonSQSConfig creates a new Config with default values.
func NewAmazonSQSConfig() AmazonSQSConfig {
	rConf := retries.NewConfig()
	rConf.Backoff.InitialInterval = "1s"
	rConf.Backoff.MaxInterval = "5s"
	rConf.Backoff.MaxElapsedTime = "30s"

	return AmazonSQSConfig{
		SessionConfig: SessionConfig{
			Config: sess.NewConfig(),
		},
		URL:                    "",
		MessageGroupID:         "",
		MessageDeduplicationID: "",
		Metadata:               metadata.NewExcludeFilterConfig(),
		MaxInFlight:            64,
		Config:                 rConf,
		Batching:               batchconfig.NewConfig(),
	}
}
