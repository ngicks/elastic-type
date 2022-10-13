package estype_test

import (
	"encoding/json"
	"testing"
	"time"
	_ "time/tzdata"

	estype "github.com/ngicks/elastic-type/es_type"
	"github.com/stretchr/testify/require"
)

var jst *time.Location

func init() {
	var err error
	jst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
}

type parseTestCase struct {
	input     any // int64 or string
	expected  time.Time
	mustError bool
}

func TestStrictDateOptionalTimeEpochMillis(t *testing.T) {
	happyPathCases := []parseTestCase{
		{
			input:    1666282966123,
			expected: time.Date(2022, 10, 20, 16, 22, 46, 123000000, time.UTC),
		},
		{
			input:    "2022-10-20T16:22:46.123+09:00",
			expected: time.Date(2022, 10, 20, 16, 22, 46, 123000000, jst),
		},
	}

	for _, testCase := range happyPathCases {
		jsonValue, _ := json.Marshal(testCase.input)
		var timeVal estype.StrictDateOptionalTimeEpochMillis
		err := json.Unmarshal(jsonValue, &timeVal)
		require.NoError(t, err)
		require.Conditionf(
			t,
			func() (success bool) { return time.Time(timeVal).Equal(testCase.expected) },
			"expected: %s, actual: %s",
			testCase.expected,
			timeVal,
		)
	}
}

// generate_date:start

func FuzzDateOptionalTime(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.DateOptionalTime(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.DateOptionalTime
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictDateOptionalTime(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictDateOptionalTime(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictDateOptionalTime
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictDateOptionalTimeNanos(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictDateOptionalTimeNanos(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictDateOptionalTimeNanos
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzBasicDate(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.BasicDate(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.BasicDate
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzBasicDateTime(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.BasicDateTime(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.BasicDateTime
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzBasicDateTimeNoMillis(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.BasicDateTimeNoMillis(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.BasicDateTimeNoMillis
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzBasicOrdinalDate(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.BasicOrdinalDate(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.BasicOrdinalDate
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzBasicOrdinalDateTime(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.BasicOrdinalDateTime(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.BasicOrdinalDateTime
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzBasicOrdinalDateTimeNoMillis(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.BasicOrdinalDateTimeNoMillis(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.BasicOrdinalDateTimeNoMillis
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzBasicTime(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.BasicTime(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.BasicTime
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzBasicTimeNoMillis(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.BasicTimeNoMillis(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.BasicTimeNoMillis
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzBasicTTime(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.BasicTTime(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.BasicTTime
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzBasicTTimeNoMillis(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.BasicTTimeNoMillis(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.BasicTTimeNoMillis
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzBasicWeekDate(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.BasicWeekDate(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.BasicWeekDate
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictBasicWeekDate(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictBasicWeekDate(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictBasicWeekDate
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzBasicWeekDateTime(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.BasicWeekDateTime(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.BasicWeekDateTime
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictBasicWeekDateTime(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictBasicWeekDateTime(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictBasicWeekDateTime
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzBasicWeekDateTimeNoMillis(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.BasicWeekDateTimeNoMillis(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.BasicWeekDateTimeNoMillis
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictBasicWeekDateTimeNoMillis(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictBasicWeekDateTimeNoMillis(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictBasicWeekDateTimeNoMillis
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzDate(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.Date(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.Date
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictDate(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictDate(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictDate
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzDateHour(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.DateHour(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.DateHour
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictDateHour(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictDateHour(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictDateHour
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzDateHourMinute(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.DateHourMinute(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.DateHourMinute
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictDateHourMinute(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictDateHourMinute(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictDateHourMinute
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzDateHourMinuteSecond(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.DateHourMinuteSecond(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.DateHourMinuteSecond
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictDateHourMinuteSecond(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictDateHourMinuteSecond(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictDateHourMinuteSecond
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzDateHourMinuteSecondFraction(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.DateHourMinuteSecondFraction(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.DateHourMinuteSecondFraction
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictDateHourMinuteSecondFraction(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictDateHourMinuteSecondFraction(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictDateHourMinuteSecondFraction
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzDateHourMinuteSecondMillis(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.DateHourMinuteSecondMillis(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.DateHourMinuteSecondMillis
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictDateHourMinuteSecondMillis(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictDateHourMinuteSecondMillis(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictDateHourMinuteSecondMillis
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzDateTime(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.DateTime(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.DateTime
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictDateTime(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictDateTime(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictDateTime
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzDateTimeNoMillis(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.DateTimeNoMillis(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.DateTimeNoMillis
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictDateTimeNoMillis(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictDateTimeNoMillis(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictDateTimeNoMillis
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzHour(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.Hour(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.Hour
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictHour(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictHour(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictHour
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzHourMinute(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.HourMinute(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.HourMinute
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictHourMinute(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictHourMinute(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictHourMinute
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzHourMinuteSecond(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.HourMinuteSecond(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.HourMinuteSecond
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictHourMinuteSecond(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictHourMinuteSecond(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictHourMinuteSecond
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzHourMinuteSecondFraction(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.HourMinuteSecondFraction(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.HourMinuteSecondFraction
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictHourMinuteSecondFraction(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictHourMinuteSecondFraction(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictHourMinuteSecondFraction
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzHourMinuteSecondMillis(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.HourMinuteSecondMillis(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.HourMinuteSecondMillis
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictHourMinuteSecondMillis(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictHourMinuteSecondMillis(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictHourMinuteSecondMillis
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzOrdinalDate(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.OrdinalDate(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.OrdinalDate
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictOrdinalDate(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictOrdinalDate(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictOrdinalDate
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzOrdinalDateTime(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.OrdinalDateTime(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.OrdinalDateTime
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictOrdinalDateTime(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictOrdinalDateTime(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictOrdinalDateTime
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzOrdinalDateTimeNoMillis(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.OrdinalDateTimeNoMillis(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.OrdinalDateTimeNoMillis
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictOrdinalDateTimeNoMillis(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictOrdinalDateTimeNoMillis(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictOrdinalDateTimeNoMillis
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzTime(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.Time(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.Time
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictTime(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictTime(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictTime
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzTimeNoMillis(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.TimeNoMillis(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.TimeNoMillis
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictTimeNoMillis(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictTimeNoMillis(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictTimeNoMillis
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzTTime(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.TTime(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.TTime
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictTTime(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictTTime(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictTTime
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzTTimeNoMillis(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.TTimeNoMillis(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.TTimeNoMillis
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictTTimeNoMillis(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictTTimeNoMillis(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictTTimeNoMillis
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzYear(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.Year(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.Year
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictYear(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictYear(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictYear
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzYearMonth(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.YearMonth(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.YearMonth
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictYearMonth(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictYearMonth(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictYearMonth
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzYearMonthDay(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.YearMonthDay(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.YearMonthDay
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

func FuzzStrictYearMonthDay(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := estype.StrictYearMonthDay(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled estype.StrictYearMonthDay
		err = json.Unmarshal(bin, &unmarshalled)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		binAgain, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		if str1, str2 := string(bin), string(binAgain); str1 != str2 {
			t.Fatalf("not equal: expected = %s, actual = %s", str1, str2)
		}
	})
}

// generate_date:end
