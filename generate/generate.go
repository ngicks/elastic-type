package generate

import "github.com/ngicks/elastic-type/mapping"

type GeneratedType struct {
	TyName  string
	TyDef   string
	Imports []string
}

type Option struct {
	IsRequired  bool
	IsSingle    bool
	ChildOption map[string]Option
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
// The Elasticsearch (or Apache Lucene) does not place any assumption that your data fields.
// opts here is meta data which will be used to build that assumption, like optional|required or single|many, by our own.
// Keys of opts must be mapping property names. ChildOption for types which can not be nested will be simply ignored.
func Generate(props mapping.Properties, opts map[string]Option) (highLevelTy, rawTy GeneratedType, err error) {
	panic("not implemented")
}
