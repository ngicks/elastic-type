package estype

import (
	"encoding/json"
	"strings"
)

// Boolean is elastic boolean type.
// It can be unmarshalled from boolean literal or string literal of "true" / "false" or "" (empty string).
// see: https://www.elastic.co/guide/en/elasticsearch/reference/8.4/boolean.html
//
// It marshals into boolean literal.
type Boolean bool

// MarshalJSON marshals this type into byte slice representing JSON boolean literal, true or false.
func (b Boolean) MarshalJSON() ([]byte, error) {
	return json.Marshal(bool(b))
}

func (b *Boolean) UnmarshalJSON(data []byte) error {
	bb, err := unmarshalEsBoolean(data)
	if err != nil {
		return err
	}
	*b = Boolean(bb)
	return nil
}

func (b Boolean) String() string {
	return stringEsBoolean(bool(b))
}

// BooleanStr is elastic boolean type.
// It can be unmarshalled from boolean literal of string literal of "true" / "false" or "" (empty string).
// see: https://www.elastic.co/guide/en/elasticsearch/reference/8.4/boolean.html
//
// It marshals into string literal, "true" or "false".
type BooleanStr bool

// MarshalJSON marshals this type into byte slice representing JSON string literal, "true" or "false".
// This never converts to an empty string.
func (b BooleanStr) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *BooleanStr) UnmarshalJSON(data []byte) error {
	bb, err := unmarshalEsBoolean(data)
	if err != nil {
		return err
	}
	*b = BooleanStr(bb)
	return nil
}

func (b BooleanStr) String() string {
	return stringEsBoolean(bool(b))
}

func stringEsBoolean(b bool) string {
	if b {
		return "true"
	} else {
		return "false"
	}
}

func unmarshalEsBoolean(data []byte) (bool, error) {
	switch strings.Trim(string(data), " ") {
	case `true`, `"true"`:
		return true, nil
	case `false`, `"false"`, `""`:
		return false, nil
	}

	return false, &InvalidTypeError{
		Type:         "Boolean",
		SupposedToBe: []any{true, false, "true", "false"},
		InputValue:   data,
	}
}
