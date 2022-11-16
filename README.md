# elastic-type

A type generator that generates Go types from [Elasticsearch](https://www.elastic.co/guide/en/elastic-stack/index.html) mappings.

## Ambition

The goal of `elastic-type` is to make it easy to

- [x] Generate Go types from externally-maintained json data.
- [x] Decode Elasticsearch `_source`s and consume it like plain Go structs.
  - Use generated type with [typed API!](https://github.com/elastic/go-elasticsearch/blob/e75332f0d382e54cd9ae2dfb3a0ef863759ad2b0/typedapi/types/document.go#L28)
- [ ] Generate strongly typed DSL-builder helpers.

## Target Elasticsearch version

It is tested only against Elasticsearch 8.4.

See [test.compose.yml](test.compose.yml).

CI settings could be used to test against many Elasticsearch versions. This is not yet planned.

## Overview: Current State

### Type generation

It generates 2 types from Elasticsearch mappings.

Raw one and high-level one, with interconversion methods.

Raw one is pretty straightforward. As its name suggests, it only exist to lossless-ly decode json data stored inside Elasticsearch instances.

Elasticsearch allows its json format to be _elastic_, where you can store keys with value of `T`, [T[] (, null[] or a nested T[] like [1, 2, 3 [4, 5]] which will be treated as flatted in the search context.)](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/array.html), `undefined` or `null`.

The raw type wraps all its field type, which is defined in your mappings.json, with estype.Field[T] to marshal / unmarshal most of (not all) those variants.

High-level one is like a plain Go struct which you define everyday. It only contains T, []T fields if your application defines them to be required, or \*T, \*[]T if they are optional. At least you will not be aware of the variants, which is mentioned earlier, with this type.

### Search DSL Helper

Planned. Not coming too soon.

### Installation

You can use exposed functions. `Generate` is a main entry point for code generation. And `WriteFile` is a write-file helper for generated types.

Or for your convenience, install and use the executable:

```
go install github.com/ngicks/elastic-type/cmd/generate-es-type@latest
```

Below, Example uses this executable.

### Example

It takes an Elasticsearch mapping as an input, and 2 additional options. For the format of options, refer to type definitions of MapOption and GlobalOption.

Use raw types to unmarshal json directly, call ToPlain on it to get high-level type structs.

```bash
generate-es-type -prefix-with-index-name -i ./example.json -out-high ./example_high.go -out-raw ./example_raw.go -global-option ./example_global_option.json -map-option ./example_map_option.json
```

```json
// example.json
// This is what you fetch from <es_origin>/<index_name>/_mappings
{
  "example": {
    "mappings": {
      "dynamic": "strict",
      "properties": {
        "blob": {
          "type": "binary"
        },
        "bool": {
          "type": "boolean"
        },
        "date": {
          "type": "date",
          "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
        }
      }
    }
  }
}
// example_global_option.json
{
  "IsSingle": true,
  "TypeOption": {
    "date": {
      "IsRequired": true,
      "IsSingle": true
    }
  }
}
// example_map_option.json
{
  "blob": {
    "IsRequired": true,
    "IsSingle": false
  }
}
```

It generates:

```go
package example

import (
	"encoding/json"
	"time"

	estype "github.com/ngicks/elastic-type/es_type"
	"github.com/ngicks/flextime"
	typeparamcommon "github.com/ngicks/type-param-common"
)

type Example struct {
	Blob [][]byte        `json:"blob"`
	Bool *estype.Boolean `json:"bool"`
	Date ExampleDate     `json:"date"`
}

func (t Example) ToRaw() ExampleRaw {
	return ExampleRaw{
		Blob: estype.NewFieldSlice(t.Blob, false),
		Bool: estype.NewFieldSinglePointer(t.Bool, false),
		Date: estype.NewFieldSingleValue(t.Date, false),
	}
}

// ExampleDate represents elasticsearch date.
type ExampleDate time.Time

func (t ExampleDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

var parserExampleDate = flextime.NewFlextime(
	typeparamcommon.Must(flextime.NewLayoutSet(`2006-01-02 15:04:05`)).
		AddLayout(typeparamcommon.Must(flextime.NewLayoutSet(`2006-01-02`))),
)

func (t *ExampleDate) UnmarshalJSON(data []byte) error {
	tt, err := estype.UnmarshalEsTime(
		data,
		parserExampleDate.Parse,
		time.UnixMilli,
	)
	if err != nil {
		return err
	}
	*t = ExampleDate(tt)
	return nil
}

func (t ExampleDate) String() string {
	return time.Time(t).Format(`2006-01-02 15:04:05`)
}
```

and

```go
package example

import (
	estype "github.com/ngicks/elastic-type/es_type"
)

type ExampleRaw struct {
	Blob estype.Field[[]byte]         `json:"blob"`
	Bool estype.Field[estype.Boolean] `json:"bool" esjson:"single"`
	Date estype.Field[ExampleDate]    `json:"date" esjson:"single"`
}

func (r ExampleRaw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t ExampleRaw) ToPlain() Example {
	return Example{
		Blob: t.Blob.ValueZero(),
		Bool: t.Bool.ValueSingle(),
		Date: t.Date.ValueSingleZero(),
	}
}
```

## packages

### es_type

Helper types for Elasticsearch type.

Elasticsearch is somehow _elastic_ for its json value.

For example:

- The boolean type can accept true/false and "true"/"false"/""(which means false).
- The date type defaults to a format of `strict_date_optional_time||epoch_millis`, which means it can accept number that represent unix milli second, or string value that is formatted as `YYYY-MM-dd'T'HH:mm:ss.S[S...(up to 9 digits)]` or `YYYY-MM-dd`.
  - This is really helpful when you define a script or an ingest pipeline, as you can store simply a long value for date.
    - Namely, `new Date().getTime()`
- The geopoint type allows 6 different notations to store.
  - The doc says it is for `historical reasons`.

Those types, which can be unmarshalled from various notations, needs user-defined UnmarshalJSON method. Every types defined in `es_type` has one.

#### MVPs

Seemingly these types need special unmarshallers.

- [x] Marshalling/Unmarshalling helper (Field[T any])
- [x] binary
- [x] boolean
- [x] date for built-in es date formats
- [ ] histogram
- [x] geopoint
- [x] geoshape
- [ ] join
- [ ] ranges
- [ ] rank_feature/rank_features
- [ ] point
  - basically same as geopoint, but fewer supported data notations.
- [ ] shape
  - basically same as geoshape.
- [ ] version

### generate

Code generator. It generates go code from an Elasticsearch mapping.

#### MVPs

- [x] Generate high level / raw types
  - The high level type is a plain go struct type, which is similar to what you define every day. That probably can not be unmarshalled from / marshalled into a json to be stored in the Elasticsearch.
  - The raw type is strictly compliant to Elasticsearch json format. All fields can be `undefined`, `null`, T or an array of T.
    - Which is achieved in help of `estype.Field[T]`
- [x] Generate raw level types marshaller code.
- [x] Generate high level / raw conversion code.
- [ ] Test using a real Elasticsearch instance.

#### Optionally we would do

- [ ] Remove overlapping type definition.
  - If two or more type definitions are exactly same, generate only one type and use it.
  - Evaluating that 2 defs are semantically same is hard without ast parser. Maybe we should do it in post-process.

### mapping

~~Type definitions for Elasticsearch mappings.~~

~~Used to parse mapping.~~

#### ~~MVPs~~

- [x] ~~Cover all mappings.~~
  - The official client had had mapping types before I started making this😭.
- [ ] Use [go-elasticsearch](https://github.dev/elastic/go-elasticsearch/tree/main/typedapi/types)
