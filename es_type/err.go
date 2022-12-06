package estype

import (
	"fmt"
)

type InvalidTypeError struct {
	// name of Go type
	Type         string
	SupposedToBe []any
	InputValue   []byte
}

func (e *InvalidTypeError) Error() string {
	return fmt.Sprintf(
		"invalid type error: input is unacceptable to type %s.\n"+
			"expected: one of %+v.\n"+
			"actual: %s",
		e.Type,
		e.SupposedToBe,
		string(e.InputValue),
	)
}
