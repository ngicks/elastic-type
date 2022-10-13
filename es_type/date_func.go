package estype

import (
	"encoding/json"
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

func UnmarshalEsTime(data []byte, strParser StrParser, numParser NumParser) (time.Time, error) {
	str := string(data)
	if strParser != nil && strings.HasPrefix(str, `"`) && strings.HasSuffix(str, `"`) {
		return strParser(str[1 : len(str)-1])
	}
	num, err := strconv.ParseInt(str, 10, 64)
	if err == nil && numParser != nil {
		return numParser(num), nil
	}

	var v any
	err = json.Unmarshal(data, &v)
	if err != nil {
		return time.Time{}, err
	}

	return time.Time{}, &InvalidTypeError{
		Type:            "StrictDateOptionalTimeEpochMillis",
		SupposedTValues: []any{"time formatted as string", "unix epoch number that convertible to int64"},
		InputValue:      v,
	}
}
