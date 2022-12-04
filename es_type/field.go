package estype

import (
	"bytes"
	"encoding/json"
)

const (
	StructTag = "esjson"
	// If a struct fields has this tag set (as `esjson:"single"`),
	// it will always marshal into single value, even if the field has many values.
	TagSingle = "single"
)

// IsEmpty determines if f would be treated as null in Elasticsearch.
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

// Field is a helper type to store an Elasticsearch JSON field.
// Field supports only undefined, null, T and T[],
// while Elasticsearch allows it to be one of undefined, null, null[], (null | T)[], T, T[] or T[][].
// see: https://www.elastic.co/guide/en/elasticsearch/reference/8.4/array.html
type Field[T any] struct {
	inner *[]T
}

// NewFieldUnsafe returns a new Field.
// This does not clone v.
func NewFieldUnsafe[T any](v *[]T) Field[T] {
	return Field[T]{
		inner: v,
	}
}

// NewField returns a new Field.
// The returned Field does not points to the data that v does.
// Instead Field will use cloned content of v if not nil.
func NewField[T any](v *[]T) Field[T] {
	if v == nil {
		var zero Field[T]
		return zero
	}
	return NewFieldSlice(*v, true)
}

// NewFieldSlice returns a new Field.
// The returned Field uses cloned v.
func NewFieldSlice[T any](v []T, nilIsNull bool) Field[T] {
	if v == nil {
		if nilIsNull {
			return NewFieldNull[T]()
		} else {
			var f Field[T]
			return f
		}
	}

	cloned := make([]T, len(v))
	copy(cloned, v)
	return Field[T]{
		inner: &cloned,
	}
}

func NewFieldSinglePointer[T any](v *T, nilIsNull bool) Field[T] {
	if v == nil {
		if nilIsNull {
			return NewFieldNull[T]()
		} else {
			var f Field[T]
			return f
		}
	}

	return Field[T]{
		inner: &[]T{*v},
	}
}

func NewFieldSingleValue[T any](v T) Field[T] {
	return Field[T]{
		inner: &[]T{v},
	}
}

func NewFieldNull[T any]() Field[T] {
	f := Field[T]{}
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

// SetEmpty sets the empty []T to f.
func (f *Field[T]) SetEmpty() {
	sl := make([]T, 0)
	f.inner = &sl
}

// SetValue sets cloned value.
func (f *Field[T]) SetValue(value []T) {
	cloned := make([]T, len(value))
	copy(cloned, value)
	f.inner = &cloned
}

func (f *Field[T]) SetSingleValue(value T) {
	f.inner = &[]T{value}
}

func (f Field[T]) Value() *[]T {
	return f.inner
}

// ValueZero gets the inner value of f, or zero value for T.
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

// ValueSingleZero gets the inner value of f, or falls back to zero value for T.
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
// ValueAny returns nil if f is null or undefined.
// For other cases:
//   - If single is true, it returns either single T or zero value of T if and only if inner value is zero length.
//   - If single is false, it returns []T. It could be length of zero.
func (f Field[T]) ValueAny(single bool) any {
	if f.IsUndefined() || f.IsNull() {
		return nil
	}

	val := f.Unwrap()
	if single {
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
// It always marshalls as []T.
//
// For most cases, a struct that contains Field[T] should be marshalled through MarshalFieldsJSON.
func (f Field[T]) MarshalJSON() ([]byte, error) {
	if f.IsUndefined() || f.IsNull() {
		return []byte("null"), nil
	}
	return json.Marshal(*f.inner)
}

func (b *Field[T]) UnmarshalJSON(data []byte) error {
	data = bytes.Trim(data, " ")
	var storedErr error
	if data[0] == '[' {
		b.SetEmpty()
		err := json.Unmarshal(data, b.inner)
		if err == nil {
			return nil
		}
		// in case of T = []U (e.g. dense_vector is []float64.)
		storedErr = err
	}

	if string(data) == "null" {
		b.SetNull()
		return nil
	}

	var single T
	err := json.Unmarshal(data, &single)
	if err != nil {
		if storedErr != nil {
			return storedErr
		} else {
			return err
		}
	}
	b.SetSingleValue(single)
	return nil
}

// MapField returns a new Field[T] whose values are elements of field mapped through mapper.
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
