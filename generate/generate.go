package generate

import "github.com/ngicks/elastic-type/mapping"

type GeneratedType struct {
	TyName  string
	TyDef   string
	Imports []string
	Option  FieldOption
}

// Generate generates Go struct types from an Elasticsearch mapping.
//
// It generates 2 type implementations, high level one and raw one.
// The high level type is like a plain go struct, where a field that should be:
//   - many, is []T
//   - single, is T
//   - optional, is *T (or *[]T)
//   - required, is T (or []T)
//
// The raw type is, on the other hand,
// a type that can be directly unmarshalled from an Elasticsearch-stored json.
// All fields in the type can be null, undefined, T or []T.
// This is done with help of estype.Field[T any].
//
// The Elasticsearch (or Apache Lucene) is so _elastic_ that you can store every above variants of T.
// opts here is meta data which will be used to build an assumption about the data format you are to store,
// like optional|required or single|many, by our own.
// Keys of opts must be mapping property names. ChildOption will be only used for Object or Nested, for other types simply ignored.
//
// Always len(highLevenTy) == len(rawTy).
func Generate(props mapping.Properties, tyName string, globalOpt GlobalOption, opts MapOption) (highLevelTy, rawTy []GeneratedType, err error) {
	return object(props, globalOpt, opts, []string{tyName})
}
