package estype

import (
	"strconv"
	"strings"
	"time"
)

type (
	StrParser = func(value string) (time.Time, error)
	NumParser = func(value int64) time.Time
)

func ParseUnixSec(v int64) time.Time {
	return time.Unix(v, 0)
}

func UnmarshalEsTime(data []byte, strParser StrParser, numParser NumParser, typeNames ...string) (time.Time, error) {
	str := strings.Trim(string(data), " ")
	if strParser != nil && strings.HasPrefix(str, `"`) && strings.HasSuffix(str, `"`) {
		return strParser(str[1 : len(str)-1])
	}
	num, err := strconv.ParseInt(str, 10, 64)
	if err == nil && numParser != nil {
		return numParser(num), nil
	}

	var typeName string
	if len(typeNames) >= 1 {
		typeName = typeNames[0]
	}
	return time.Time{}, &InvalidTypeError{
		Type:         typeName,
		SupposedToBe: []any{"time formatted as string", "unix epoch number that convertible to int64"},
		InputValue:   data,
	}
}
