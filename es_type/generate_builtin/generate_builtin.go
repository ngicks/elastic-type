package main

import (
	"bytes"
	"flag"
	"html/template"
	"io"
	"os"
	"strings"

	"github.com/ngicks/type-param-common/slice"
)

const codeInsertionPoint = "// generate_date"

var (
	outputFile     = flag.String("out", "", "")
	outputTestFile = flag.String("test", "", "")
)

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func main() {
	flag.Parse()

	typ := make([]string, 0)
	fuzzTest := make([]string, 0)
	for _, formatName := range target {
		buf := bytes.NewBuffer(make([]byte, 0))

		err := tyTmpl.Execute(buf, templateParam{TyName: formatName, FormatName: formatName})
		if err != nil {
			panic(err)
		}
		typ = append(typ, buf.String())

		buf.Reset()
		err = testTmpl.Execute(buf, testTmplParam{TyName: formatName, PackageName: "estype"})
		if err != nil {
			panic(err)
		}
		fuzzTest = append(fuzzTest, buf.String())
	}

	var file *os.File

	file = must(os.OpenFile(*outputFile, os.O_RDWR, 0o666))
	typeAdded := replaceBetweenComment(file, codeInsertionPoint, typ)
	file = must(os.Create(*outputFile))
	must(io.WriteString(file, typeAdded))

	file = must(os.OpenFile(*outputTestFile, os.O_RDWR, 0o666))
	testAdded := replaceBetweenComment(file, codeInsertionPoint, fuzzTest)
	file = must(os.Create(*outputTestFile))
	must(io.WriteString(file, testAdded))
}

func replaceBetweenComment(file *os.File, codeInsertionPoint string, inserted []string) string {
	text := string(must(io.ReadAll(file)))
	lines := strings.Split(text, "\n")

	start := slice.Position(lines, func(v string) bool { return strings.HasPrefix(v, codeInsertionPoint) })
	end := slice.PositionLast(lines, func(v string) bool { return strings.HasPrefix(v, codeInsertionPoint) })

	if start < 0 || end < 0 || start == end {
		panic("code insertion point not found: " + codeInsertionPoint)
	}

	out := append(append(lines[:start+1], inserted...), lines[end:]...)

	return strings.Join(out, "\n")
}

type templateParam struct {
	TyName     string
	FormatName string
}

var tyTmpl = template.Must(template.New("v").Parse(`
type {{.TyName}} time.Time

func (t {{.TyName}}) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *{{.TyName}}) UnmarshalJSON(data []byte) error {
	tt, err := UnmarshalEsTime(
		data,
		builtinformat.Formatters[builtinformat.{{.FormatName}}].Parse,
		nil,
		` + "`" + `{{.TyName}}` + "`" + `,
	)
	if err != nil {
		return err
	}
	*t = {{.TyName}}(tt)
	return nil
}

func (t {{.TyName}}) String() string {
	return time.Time(t).Format(builtinformat.FormatLayout[builtinformat.{{.FormatName}}])
}
`))

type testTmplParam struct {
	TyName      string
	PackageName string
}

var testTmpl = template.Must(template.New("n").Parse(`
func Fuzz{{.TyName}}(f *testing.F) {
	f.Add(int64(1666282966123), int64(218964189023))
	f.Fuzz(func(t *testing.T, milliSec int64, nanoSec int64) {
		tt := {{.PackageName}}.{{.TyName}}(time.UnixMilli(milliSec).Add(time.Duration(nanoSec)))

		bin, err := json.Marshal(tt)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}
		var unmarshalled {{.PackageName}}.{{.TyName}}
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
`))

var target = [...]string{
	"DateOptionalTime",
	"StrictDateOptionalTime",
	"StrictDateOptionalTimeNanos",
	"BasicDate",
	"BasicDateTime",
	"BasicDateTimeNoMillis",
	"BasicOrdinalDate",
	"BasicOrdinalDateTime",
	"BasicOrdinalDateTimeNoMillis",
	"BasicTime",
	"BasicTimeNoMillis",
	"BasicTTime",
	"BasicTTimeNoMillis",
	"BasicWeekDate",
	"StrictBasicWeekDate",
	"BasicWeekDateTime",
	"StrictBasicWeekDateTime",
	"BasicWeekDateTimeNoMillis",
	"StrictBasicWeekDateTimeNoMillis",
	"Date",
	"StrictDate",
	"DateHour",
	"StrictDateHour",
	"DateHourMinute",
	"StrictDateHourMinute",
	"DateHourMinuteSecond",
	"StrictDateHourMinuteSecond",
	"DateHourMinuteSecondFraction",
	"StrictDateHourMinuteSecondFraction",
	"DateHourMinuteSecondMillis",
	"StrictDateHourMinuteSecondMillis",
	"DateTime",
	"StrictDateTime",
	"DateTimeNoMillis",
	"StrictDateTimeNoMillis",
	"Hour",
	"StrictHour",
	"HourMinute",
	"StrictHourMinute",
	"HourMinuteSecond",
	"StrictHourMinuteSecond",
	"HourMinuteSecondFraction",
	"StrictHourMinuteSecondFraction",
	"HourMinuteSecondMillis",
	"StrictHourMinuteSecondMillis",
	"OrdinalDate",
	"StrictOrdinalDate",
	"OrdinalDateTime",
	"StrictOrdinalDateTime",
	"OrdinalDateTimeNoMillis",
	"StrictOrdinalDateTimeNoMillis",
	"Time",
	"StrictTime",
	"TimeNoMillis",
	"StrictTimeNoMillis",
	"TTime",
	"StrictTTime",
	"TTimeNoMillis",
	"StrictTTimeNoMillis",
	"Year",
	"StrictYear",
	"YearMonth",
	"StrictYearMonth",
	"YearMonthDay",
	"StrictYearMonthDay",
}
