package estype

import (
	"bytes"
	"encoding/json"
)

const (
	StructTag = "esjson"
	// If set to struct tag (`esjson:"single"`),
	// marshal into always single value, even if the field has many values.
	TagSingle = "single"
)

// IsEmpty determines if field would be treated as empty in Elastichsearch.
// In its search context, null field is one of null, undefined (nonexistent), or an empty array.
func IsEmpty[T any](val *[]T) bool {
	if val == nil || len(*val) == 0 {
		// len of nil slice is 0 anyway.
		return true
	}
	return false
}

func IsNull[T any](val *[]T) bool {
	return val != nil && *val == nil
}

func IsUndefined[T any](val *[]T) bool {
	return val == nil
}

func UnwrapValue[T any](val *[]T) []T {
	if *val == nil {
		panic("nil slice")
	}
	return *val
}

// Field is an Elasticsearch field helper type.
// A Field value can be null, undefined, T or an array of T.
// It also can be a nested array but is not supported by this struct.
// see: https://www.elastic.co/guide/en/elasticsearch/reference/8.4/array.html
type Field[T any] struct {
	inner *[]T
	// If true, it marshals into an array even when its inner value slice is of single element.
	ShouldRetainArray bool
}

func NewField[T any](v *[]T, shouldRetainArray bool) Field[T] {
	return Field[T]{
		inner:             v,
		ShouldRetainArray: shouldRetainArray,
	}
}

func NewFieldSlice[T any](v []T, shouldRetainArray bool) Field[T] {
	return Field[T]{
		inner:             &v,
		ShouldRetainArray: shouldRetainArray,
	}
}

func NewFieldSinglePointer[T any](v *T, shouldRetainArray bool) Field[T] {
	if v != nil {
		return Field[T]{
			inner:             &[]T{*v},
			ShouldRetainArray: shouldRetainArray,
		}
	}
	return NewFieldNull[T](shouldRetainArray)
}

func NewFieldSingleValue[T any](v T, shouldRetainArray bool) Field[T] {
	return Field[T]{
		inner:             &[]T{v},
		ShouldRetainArray: shouldRetainArray,
	}
}

func NewFieldNull[T any](shouldRetainArray bool) Field[T] {
	f := Field[T]{
		ShouldRetainArray: shouldRetainArray,
	}
	f.SetNull()
	return f
}

func (f Field[T]) IsNull() bool {
	return IsNull(f.inner)
}

func (f Field[T]) IsUndefined() bool {
	return IsUndefined(f.inner)
}

func (f Field[T]) IsEmpty() bool {
	return IsEmpty(f.inner)
}

func (f *Field[T]) SetNull() {
	var typedNil []T
	f.inner = &typedNil
}

func (f *Field[T]) SetUndefined() {
	f.inner = nil
}

func (f *Field[T]) SetValue(value []T) {
	f.inner = &value
}

func (f *Field[T]) SetSingleValue(value T) {
	f.inner = &[]T{value}
}

func (f Field[T]) Value() *[]T {
	return f.inner
}

// ValueZero gets the inner value of f or zero value for T.
// A returned value must be non-nil []T.
// If the inner value is non-empty, it returns that value.
// Otherwise, returns newly created zero-length []T.
func (f Field[T]) ValueZero() []T {
	v := f.Value()
	if v == nil || len(*v) == 0 {
		return []T{}
	}
	return *v
}

func (f Field[T]) ValueSingle() *T {
	if f.inner == nil {
		return nil
	}
	if len(*f.inner) > 0 {
		return &(*f.inner)[0]
	}
	return nil
}

// ValueSingleZero() gets the inner value of f or zero value for T.
// If the inner value non-empty, it returns that value.
// Otherwise, returns zero value of T.
func (f Field[T]) ValueSingleZero() T {
	if v := f.ValueSingle(); v != nil {
		return *v
	} else {
		var zero T
		return zero
	}
}

// ValueAny returns inner value in any type.
// This can be used without any instantiation.
//
// If mustSingle is true, value can be a single T,
// or if mustSingle is true and the inner value is empty []T,
// return zero value of T.
func (f Field[T]) ValueAny(mustSingle bool) any {
	if f.IsUndefined() || f.IsNull() {
		return nil
	}

	val := f.Unwrap()
	if mustSingle {
		if len(val) == 0 {
			var zero T
			return zero
		} else {
			return val[0]
		}
	} else {
		return val
	}
}

func (f Field[T]) Unwrap() []T {
	return UnwrapValue(f.inner)
}

func (f Field[T]) UnwrapSingle() T {
	return (*f.inner)[0]
}

// MarshalJSON encodes f into a json format.
//
// If f only has a single element and ShouldRetainArray is false, T will be marshalled.
// It marshalls as []T otherwise.
//
// For most cases, a struct that only contains Field[T] should be marshalled through MarshalFieldsJSON.
func (f Field[T]) MarshalJSON() ([]byte, error) {
	if f.IsUndefined() || f.IsNull() {
		return []byte("null"), nil
	}
	if len(*f.inner) == 1 && !f.ShouldRetainArray {
		return json.Marshal(f.UnwrapSingle())
	}
	return json.Marshal(*f.inner)
}

func (b *Field[T]) UnmarshalJSON(data []byte) error {
	data = bytes.Trim(data, " ")
	if data[0] == '[' {
		return json.Unmarshal(data, b.inner)
	}

	if string(data) == "null" {
		b.SetNull()
		return nil
	}

	var single T
	err := json.Unmarshal(data, &single)
	if err != nil {
		return err
	}
	b.SetSingleValue(single)
	return nil
}

func MapField[T, U any](field Field[T], mapper func(v T) U) Field[U] {
	var f Field[U]
	if field.IsUndefined() {
		return f
	}
	if field.IsNull() {
		f.SetNull()
		return f
	}

	var newVal []U
	for _, v := range field.Unwrap() {
		newVal = append(newVal, mapper(v))
	}
	f.SetValue(newVal)
	return f
}
