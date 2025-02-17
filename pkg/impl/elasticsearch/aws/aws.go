package aws

import (
	"errors"

	"github.com/olivere/elastic/v7"
	aws "github.com/olivere/elastic/v7/aws/v4"

	baws "github.com/nehal119/benthos-119/pkg/impl/aws"
	"github.com/nehal119/benthos-119/pkg/impl/elasticsearch"
	"github.com/nehal119/benthos-119/public/service"
)

func init() {
	elasticsearch.AWSOptFn = func(conf *service.ParsedConfig) ([]elastic.ClientOptionFunc, error) {
		if enabled, _ := conf.FieldBool(elasticsearch.ESOFieldAWSEnabled); !enabled {
			return nil, nil
		}

		tsess, err := baws.GetSession(conf)
		if err != nil {
			return nil, err
		}

		var region string
		if tsess.Config.Region != nil {
			region = *tsess.Config.Region
		} else {
			return nil, errors.New("unable to detect target AWS region, if you encounter this error please report it via: https://github.com/benthosdev/benthos/issues/new")
		}

		signingClient := aws.NewV4SigningClient(tsess.Config.Credentials, region)
		return []elastic.ClientOptionFunc{elastic.SetHttpClient(signingClient)}, nil
	}
}
