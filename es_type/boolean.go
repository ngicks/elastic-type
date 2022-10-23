package estype

import (
	"encoding/json"
)

type Boolean bool

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

type BooleanStr bool

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
	switch string(data) {
	case `true`, `"true"`:
		return true, nil
	case `false`, `"false"`, `""`:
		return false, nil
	}

	var v any
	err := json.Unmarshal(data, &v)
	if err != nil {
		// this should not happen. validity is checked before it reaches this line.
		return false, err
	}

	return false, &InvalidTypeError{
		Type:            "Boolean",
		SupposedTValues: []any{true, false, "true", "false"},
		InputValue:      v,
	}
}
