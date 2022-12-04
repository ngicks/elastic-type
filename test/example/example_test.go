package example

import (
	"encoding/json"
	"testing"
	"time"
)

func FuzzExampleDate(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964089023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := ExampleDate(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled ExampleDate
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
