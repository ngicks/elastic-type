package example

import (
	"encoding/json"
	"testing"
	"time"
)

func FuzzAllDate(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964089023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := AllDate(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled AllDate
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

func FuzzAllDateNano(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964089023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := AllDateNano(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled AllDateNano
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
