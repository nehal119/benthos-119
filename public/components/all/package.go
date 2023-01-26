// Package all imports all component implementations that ship with the open
// source Benthos repo. This is a convenient way of importing every single
// connector at the cost of a larger dependency tree for your application.
package all

import (
	// Import all public sub-categories.
	_ "github.com/nehal119/benthos-119/public/components/amqp09"
	_ "github.com/nehal119/benthos-119/public/components/amqp1"
	_ "github.com/nehal119/benthos-119/public/components/avro"
	_ "github.com/nehal119/benthos-119/public/components/aws"
	_ "github.com/nehal119/benthos-119/public/components/azure"
	_ "github.com/nehal119/benthos-119/public/components/beanstalkd"
	_ "github.com/nehal119/benthos-119/public/components/cassandra"
	_ "github.com/nehal119/benthos-119/public/components/confluent"
	_ "github.com/nehal119/benthos-119/public/components/couchbase"
	_ "github.com/nehal119/benthos-119/public/components/dgraph"
	_ "github.com/nehal119/benthos-119/public/components/elasticsearch"
	_ "github.com/nehal119/benthos-119/public/components/gcp"
	_ "github.com/nehal119/benthos-119/public/components/hdfs"
	_ "github.com/nehal119/benthos-119/public/components/influxdb"
	_ "github.com/nehal119/benthos-119/public/components/io"
	_ "github.com/nehal119/benthos-119/public/components/jaeger"
	_ "github.com/nehal119/benthos-119/public/components/kafka"
	_ "github.com/nehal119/benthos-119/public/components/maxmind"
	_ "github.com/nehal119/benthos-119/public/components/memcached"
	_ "github.com/nehal119/benthos-119/public/components/mongodb"
	_ "github.com/nehal119/benthos-119/public/components/mqtt"
	_ "github.com/nehal119/benthos-119/public/components/nanomsg"
	_ "github.com/nehal119/benthos-119/public/components/nats"
	_ "github.com/nehal119/benthos-119/public/components/nsq"
	_ "github.com/nehal119/benthos-119/public/components/otlp"
	_ "github.com/nehal119/benthos-119/public/components/prometheus"
	_ "github.com/nehal119/benthos-119/public/components/pure"
	_ "github.com/nehal119/benthos-119/public/components/pure/extended"
	_ "github.com/nehal119/benthos-119/public/components/pusher"
	_ "github.com/nehal119/benthos-119/public/components/redis"
	_ "github.com/nehal119/benthos-119/public/components/sftp"
	_ "github.com/nehal119/benthos-119/public/components/snowflake"
	_ "github.com/nehal119/benthos-119/public/components/sql"
	_ "github.com/nehal119/benthos-119/public/components/statsd"
	_ "github.com/nehal119/benthos-119/public/components/wasm"
)
