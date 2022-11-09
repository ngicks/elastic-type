package estype

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"

	"github.com/ngicks/type-param-common/iterator"
	"github.com/ngicks/type-param-common/slice"
)

var ErrIncorrectType = errors.New("incorrect")

// Parts of Field[T] that can be used without an instantiation.
type UninstantiatedField interface {
	ValueAny(mustSingle bool) any
	IsNull() bool
	IsUndefined() bool
}

// MarshalFieldsJSON encodes v into JSON.
// All fields of v must be Field[T any].
// If a field is set as null, outputs []byte(`null`), if set as undefined, it skips the field.
//
// v must:
//   - be a struct type
//   - have no unexported fields
//   - have only Field[T] fields.
//
// If v:
//   - is not struct, returns ErrIncorrectType.
//   - has any fields whose type is not Field, returns ErrIncorrectType.
//   - has any unexported field, then its behavior is undefined.
//
// MarshalFieldsJSON retrieves underlying values of Field type by calling ValueAny.
// Then value will be unmarshalled with json.Unmarshal. If unmarshalling returns err,
// then MarshalFieldsJSON returns the error.
func MarshalFieldsJSON(v any) ([]byte, error) {
	out := []byte(`{`)

	rv := reflect.ValueOf(v)
	rt := rv.Type()

	if rt.Kind() != reflect.Struct {
		return nil, ErrIncorrectType
	}

	for i := 0; i < rv.NumField(); i++ {
		value, ok := rv.Field(i).Interface().(UninstantiatedField)
		if !ok {
			return nil, ErrIncorrectType
		}

		if value.IsUndefined() {
			// skip this field.
			continue
		}

		field := rt.Field(i)

		name := getFieldName(field)
		out = append(out, []byte(`"`+name+`":`)...)

		if value.IsNull() {
			out = append(out, []byte("null,")...)
			continue
		}

		esFieldTags := getTag(field.Tag, StructTag)
		mustSingle := slice.Has(esFieldTags, TagSingle)

		val := value.ValueAny(mustSingle)

		encoded, err := json.Marshal(val)
		if err != nil {
			return nil, err
		}

		out = append(out, encoded...)
		out = append(out, ',')
	}

	// removing last ','
	if out[len(out)-1] == ',' {
		out = out[:len(out)-1]
	}
	out = append(out, '}')

	return out, nil
}

func getTag(tag reflect.StructTag, tagName string) []string {
	return iterator.FromSlice(
		strings.Split(
			tag.Get(tagName),
			",",
		)).
		Map(strings.TrimSpace).
		Exclude(func(s string) bool { return s == "" }).
		Collect()
}

func getFieldName(field reflect.StructField) string {
	tags := getTag(field.Tag, "json")
	if len(tags) > 0 {
		return tags[0]
	} else {
		return field.Name
	}
}
