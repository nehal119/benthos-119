// Package extended contains component implementations that have a larger
// dependency footprint but do not interact with external systems (so an
// extension of pure components)
//
// EXPERIMENTAL: The specific components excluded by this package may change
// outside of major version releases. This means we may choose to remove certain
// plugins if we determine that their dependencies are likely to interfere with
// the goals of this package.
package extended

import (
	// Import pure but larger packages.
	_ "github.com/nehal119/benthos-119/pkg/impl/awk"
	_ "github.com/nehal119/benthos-119/pkg/impl/jsonpath"
	_ "github.com/nehal119/benthos-119/pkg/impl/lang"
	_ "github.com/nehal119/benthos-119/pkg/impl/msgpack"
	_ "github.com/nehal119/benthos-119/pkg/impl/parquet"
	_ "github.com/nehal119/benthos-119/pkg/impl/pure/extended"
	_ "github.com/nehal119/benthos-119/pkg/impl/xml"
)
