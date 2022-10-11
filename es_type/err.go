package estype

import (
	"fmt"
	"reflect"
)

type InvalidTypeError struct {
	// name of Go type
	Type            string
	SupposedTValues []any
	InputValue      any
}

func (e *InvalidTypeError) Error() string {
	return fmt.Sprintf(
		"invalid type error: input is unacceptable to type %s.\n"+
			"expected: one of %+v.\n"+
			"actual: type of %s, value of %+v",
		e.Type,
		e.SupposedTValues,
		reflect.TypeOf(e.InputValue).Kind(),
		e.InputValue,
	)
}
