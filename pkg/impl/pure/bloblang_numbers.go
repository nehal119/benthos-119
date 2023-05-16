package pure

import (
	"github.com/nehal119/benthos-119/pkg/bloblang/query"
	"github.com/nehal119/benthos-119/public/bloblang"
)

func init() {
	if err := bloblang.RegisterMethodV2("int64",
		bloblang.NewPluginSpec().
			Category(query.MethodCategoryNumbers).
			Description(`
Converts a numerical type into a 64-bit signed integer, this is for advanced use cases where a specific data type is needed for a given component (such as the ClickHouse SQL driver).

If the value is a string then an attempt will be made to parse it as a 64-bit integer. If the target value exceeds the capacity of an integer or contains decimal values then this method will throw an error. In order to convert a floating point number containing decimals first use `+"[`.round()`](#round)"+` on the value first. Please refer to the [`+"`strconv.ParseInt`"+` documentation](https://pkg.go.dev/strconv#ParseInt) for details regarding the supported formats.`).
			Example("", `
root.a = this.a.int64()
root.b = this.b.round().int64()
root.c = this.c.int64()
`,
				[2]string{
					`{"a":12,"b":12.34,"c":"12"}`,
					`{"a":12,"b":12,"c":12}`,
				},
			).
			Example("", `
root = this.int64()
`,
				[2]string{
					`"0xDEADBEEF"`,
					`3735928559`,
				},
			),
		func(args *bloblang.ParsedParams) (bloblang.Method, error) {
			return func(input any) (any, error) {
				return query.IToInt(input)
			}, nil
		}); err != nil {
		panic(err)
	}

	if err := bloblang.RegisterMethodV2("int32",
		bloblang.NewPluginSpec().
			Category(query.MethodCategoryNumbers).
			Description(`
Converts a numerical type into a 32-bit signed integer, this is for advanced use cases where a specific data type is needed for a given component (such as the ClickHouse SQL driver).

If the value is a string then an attempt will be made to parse it as a 32-bit integer. If the target value exceeds the capacity of an integer or contains decimal values then this method will throw an error. In order to convert a floating point number containing decimals first use `+"[`.round()`](#round)"+` on the value first. Please refer to the [`+"`strconv.ParseInt`"+` documentation](https://pkg.go.dev/strconv#ParseInt) for details regarding the supported formats.`).
			Example("", `
root.a = this.a.int32()
root.b = this.b.round().int32()
root.c = this.c.int32()
`,
				[2]string{
					`{"a":12,"b":12.34,"c":"12"}`,
					`{"a":12,"b":12,"c":12}`,
				},
			).
			Example("", `
root = this.int32()
`,
				[2]string{
					`"0xB70B"`,
					`46859`,
				},
			),
		func(args *bloblang.ParsedParams) (bloblang.Method, error) {
			return func(input any) (any, error) {
				return query.IToInt32(input)
			}, nil
		}); err != nil {
		panic(err)
	}

	if err := bloblang.RegisterMethodV2("uint64",
		bloblang.NewPluginSpec().
			Category(query.MethodCategoryNumbers).
			Description(`
Converts a numerical type into a 64-bit unsigned integer, this is for advanced use cases where a specific data type is needed for a given component (such as the ClickHouse SQL driver).

If the value is a string then an attempt will be made to parse it as a 64-bit unsigned integer. If the target value exceeds the capacity of an integer or contains decimal values then this method will throw an error. In order to convert a floating point number containing decimals first use `+"[`.round()`](#round)"+` on the value first. Please refer to the [`+"`strconv.ParseInt`"+` documentation](https://pkg.go.dev/strconv#ParseInt) for details regarding the supported formats.`).
			Example("", `
root.a = this.a.uint64()
root.b = this.b.round().uint64()
root.c = this.c.uint64()
root.d = this.d.uint64().catch(0)
`,
				[2]string{
					`{"a":12,"b":12.34,"c":"12","d":-12}`,
					`{"a":12,"b":12,"c":12,"d":0}`,
				},
			).
			Example("", `
root = this.uint64()
`,
				[2]string{
					`"0xDEADBEEF"`,
					`3735928559`,
				},
			),
		func(args *bloblang.ParsedParams) (bloblang.Method, error) {
			return func(input any) (any, error) {
				return query.IToUint(input)
			}, nil
		}); err != nil {
		panic(err)
	}

	if err := bloblang.RegisterMethodV2("uint32",
		bloblang.NewPluginSpec().
			Category(query.MethodCategoryNumbers).
			Description(`
Converts a numerical type into a 32-bit unsigned integer, this is for advanced use cases where a specific data type is needed for a given component (such as the ClickHouse SQL driver).

If the value is a string then an attempt will be made to parse it as a 32-bit unsigned integer. If the target value exceeds the capacity of an integer or contains decimal values then this method will throw an error. In order to convert a floating point number containing decimals first use `+"[`.round()`](#round)"+` on the value first. Please refer to the [`+"`strconv.ParseInt`"+` documentation](https://pkg.go.dev/strconv#ParseInt) for details regarding the supported formats.`).
			Example("", `
root.a = this.a.uint32()
root.b = this.b.round().uint32()
root.c = this.c.uint32()
root.d = this.d.uint32().catch(0)
`,
				[2]string{
					`{"a":12,"b":12.34,"c":"12","d":-12}`,
					`{"a":12,"b":12,"c":12,"d":0}`,
				},
			).
			Example("", `
root = this.uint32()
`,
				[2]string{
					`"0xB70B"`,
					`46859`,
				},
			),
		func(args *bloblang.ParsedParams) (bloblang.Method, error) {
			return func(input any) (any, error) {
				return query.IToUint32(input)
			}, nil
		}); err != nil {
		panic(err)
	}

	if err := bloblang.RegisterMethodV2("uint8",
		bloblang.NewPluginSpec().
			Category(query.MethodCategoryNumbers).
			Description(`
Converts a numerical type into a 8-bit unsigned integer, this is for advanced use cases where a specific data type is needed for a given component (such as the ClickHouse SQL driver).

If the value is a string then an attempt will be made to parse it as a 32-bit unsigned integer. If the target value exceeds the capacity of an integer or contains decimal values then this method will throw an error. In order to convert a floating point number containing decimals first use `+"[`.round()`](#round)"+` on the value first.`).
			Example("", `
root.a = this.a.uint8()
root.b = this.b.round().uint8()
root.c = this.c.uint8()
root.d = this.d.uint8().catch(0)
`,
				[2]string{
					`{"a":12,"b":12.34,"c":"12","d":-12}`,
					`{"a":12,"b":12,"c":12,"d":0}`,
				},
			),
		func(args *bloblang.ParsedParams) (bloblang.Method, error) {
			return func(input any) (any, error) {
				return query.IToUint8(input)
			}, nil
		}); err != nil {
		panic(err)
	}

	if err := bloblang.RegisterMethodV2("uint16",
		bloblang.NewPluginSpec().
			Category(query.MethodCategoryNumbers).
			Description(`
Converts a numerical type into a 16-bit unsigned integer, this is for advanced use cases where a specific data type is needed for a given component (such as the ClickHouse SQL driver).

If the value is a string then an attempt will be made to parse it as a 32-bit unsigned integer. If the target value exceeds the capacity of an integer or contains decimal values then this method will throw an error. In order to convert a floating point number containing decimals first use `+"[`.round()`](#round)"+` on the value first.`).
			Example("", `
root.a = this.a.uint16()
root.b = this.b.round().uint16()
root.c = this.c.uint16()
root.d = this.d.uint16().catch(0)
`,
				[2]string{
					`{"a":12,"b":12.34,"c":"12","d":-12}`,
					`{"a":12,"b":12,"c":12,"d":0}`,
				},
			),
		func(args *bloblang.ParsedParams) (bloblang.Method, error) {
			return func(input any) (any, error) {
				return query.IToUint32(input)
			}, nil
		}); err != nil {
		panic(err)
	}

	if err := bloblang.RegisterMethodV2("int8",
		bloblang.NewPluginSpec().
			Category(query.MethodCategoryNumbers).
			Description(`
Converts a numerical type into a 8-bit signed integer, this is for advanced use cases where a specific data type is needed for a given component (such as the ClickHouse SQL driver).

If the value is a string then an attempt will be made to parse it as a 8-bit integer. If the target value exceeds the capacity of an integer or contains decimal values then this method will throw an error. In order to convert a floating point number containing decimals first use `+"[`.round()`](#round)"+` on the value first.`).
			Example("", `
root.a = this.a.int8()
root.b = this.b.round().int8()
root.c = this.c.int8()
`,
				[2]string{
					`{"a":12,"b":12.34,"c":"12"}`,
					`{"a":12,"b":12,"c":12}`,
				},
			),
		func(args *bloblang.ParsedParams) (bloblang.Method, error) {
			return func(input any) (any, error) {
				return query.IToInt16(input)
			}, nil
		}); err != nil {
		panic(err)
	}

	if err := bloblang.RegisterMethodV2("int16",
		bloblang.NewPluginSpec().
			Category(query.MethodCategoryNumbers).
			Description(`
Converts a numerical type into a 16-bit signed integer, this is for advanced use cases where a specific data type is needed for a given component (such as the ClickHouse SQL driver).

If the value is a string then an attempt will be made to parse it as a 8-bit integer. If the target value exceeds the capacity of an integer or contains decimal values then this method will throw an error. In order to convert a floating point number containing decimals first use `+"[`.round()`](#round)"+` on the value first.`).
			Example("", `
root.a = this.a.int16()
root.b = this.b.round().int16()
root.c = this.c.int16()
`,
				[2]string{
					`{"a":12,"b":12.34,"c":"12"}`,
					`{"a":12,"b":12,"c":12}`,
				},
			),
		func(args *bloblang.ParsedParams) (bloblang.Method, error) {
			return func(input any) (any, error) {
				return query.IToInt16(input)
			}, nil
		}); err != nil {
		panic(err)
	}
}
