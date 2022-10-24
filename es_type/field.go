package estype

import "encoding/json"

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

// Field is an Elastichsearch field helper type.
// A Field value can be null, undefined, T or an array of T.
// It also can be a nested array but is not supported by this struct.
// see: https://www.elastic.co/guide/en/elasticsearch/reference/8.4/array.html
type Field[T any] struct {
	inner *[]T
	// If true, it marshals into an array even when its inner value slice is of single element.
	ShouldRetainArray bool
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

func (f Field[T]) ValueSingle() *T {
	if f.inner == nil {
		return nil
	}
	if len(*f.inner) > 0 {
		return &(*f.inner)[0]
	}
	return nil
}

func (f Field[T]) Unwrap() []T {
	return UnwrapValue(f.inner)
}

func (f Field[T]) UnwrapSingle() T {
	return (*f.inner)[0]
}

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
