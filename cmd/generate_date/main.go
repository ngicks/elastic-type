package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	builtinformat "github.com/ngicks/elastic-type/es_type/builtin_format"
	"github.com/ngicks/flextime"
	"github.com/ngicks/type-param-common/set"
	"github.com/ngicks/type-param-common/slice"
)

var (
	typeName           = flag.String("type-name", "EsDate", "type name to generate.")
	formatStr          = flag.String("formats", "", "[required]formats of date, double vertical line(`||`) splitted. e.g. yyyy-MM-dd'T'HH:mm:ss.SSSZ||epoch_millis.")
	out                = flag.String("out", "", "output filename. stdout if not set.")
	includeImports     = flag.Bool("include-imports", false, "if true, output includes import directives.")
	preferEpochMarshal = flag.Bool("prefer-epoch", false, "if true, ouput type mashals into epoch milli or second.")
	preferdFormat      = flag.String(
		"prefered-format",
		"",
		"prefered format to use with MarshalJSON and String method. "+
			"If empty, generated code uses longest string format. "+
			"Value must be one of formats. -prefer-epoch has precedence.",
	)
)

func main() {
	flag.Parse()

	if *formatStr == "" {
		panic("required cli arg not given")
	}

	formats := strings.Split(*formatStr, "||")
	safeFormats := make([]string, 0)
	hasEpochFormat := false
	isEpochMillis := false
	formatSet := set.New[string]()
	for _, format := range formats {
		formatSet.Add(format)
		v, ok := builtinformat.ParsingLayout[format]
		if ok {
			safeFormats = append(safeFormats, v)
		} else if format == builtinformat.EpochMillis || format == builtinformat.EpochSecond {
			if hasEpochFormat {
				panic("two or more epoch format")
			}
			hasEpochFormat = true
			if format == builtinformat.EpochMillis {
				isEpochMillis = true
			}
		} else {
			safeFormats = append(safeFormats, format)
		}
	}

	if formatSet.Len() != len(formats) {
		panic("dupe entry in date formats")
	}

	layoutset := preflightParse(safeFormats)

	var marshallingFormat string
	if *preferdFormat == "" {
		marshallingFormat = layoutset.Layout()[0]
	} else {
		converted, err := flextime.ReplaceTimeToken(*preferdFormat)
		if err != nil {
			panic(err)
		}
		if !slice.Has(layoutset.Layout(), converted) {
			panic(fmt.Sprintf("prefered formated %s is not one of formats %+v", *preferdFormat, layoutset.Layout()))
		}
		marshallingFormat = converted
	}

	var outputW io.Writer
	if *out == "" {
		outputW = os.Stdout
	} else {
		f, err := os.Create(*out)
		if err != nil {
			panic(err)
		}
		outputW = f
		defer f.Close()
	}

	if *includeImports {
		err := importTmpl.Execute(outputW, struct{ PreferEpoch bool }{PreferEpoch: *preferEpochMarshal})
		if err != nil {
			panic(err)
		}
	}

	err := tyTmpl.Execute(outputW, tyParam{
		TyName:            *typeName,
		StrFormats:        safeFormats,
		MarshallingFormat: marshallingFormat,
		HasNumFormat:      hasEpochFormat,
		NumFormatIsMillis: isEpochMillis,
		PreferEpoch:       *preferEpochMarshal,
	})
	if err != nil {
		panic(err)
	}
}

func preflightParse(formats []string) *flextime.LayoutSet {
	first, formats := formats[0], formats[1:]

	layoutSet, err := flextime.NewLayoutSet(first)
	if err != nil {
		panic(err)
	}
	for _, format := range formats {
		additive, err := flextime.NewLayoutSet(format)
		if err != nil {
			panic(err)
		}
		layoutSet = layoutSet.AddLayout(additive)
	}
	return layoutSet
}

var importTmpl = template.Must(template.New("n").Parse(`import (
	"encoding/json"
	{{- if .PreferEpoch}}
	"strconv"
	{{- end}}
	"time"

	estype "github.com/ngicks/elastic-type/es_type"
	"github.com/ngicks/flextime"
	typeparamcommon "github.com/ngicks/type-param-common"
)`))

type tyParam struct {
	TyName            string
	StrFormats        []string
	MarshallingFormat string
	HasNumFormat      bool
	NumFormatIsMillis bool
	PreferEpoch       bool
}

var tyTmpl = template.Must(template.New("v").Parse(`
// {{.TyName}} represents elasticsearch date.
type {{.TyName}} time.Time

func (t {{.TyName}}) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}


var parser{{.TyName}} = flextime.NewFlextime(
{{range $index, $format := .StrFormats}}
	{{- if eq $index 0 -}}
	typeparamcommon.Must(flextime.NewLayoutSet(` + "`" + `{{$format}}` + "`" + `))
	{{- else -}}
	.
	AddLayout(typeparamcommon.Must(flextime.NewLayoutSet(` + "`" + `{{$format}}` + "`" + `)))
	{{- end -}}
{{- end -}},
)

func (t *{{.TyName}}) UnmarshalJSON(data []byte) error {
	tt, err := estype.UnmarshalEsTime(
		data, 
		parser{{.TyName}}.Parse, 
	{{- if .HasNumFormat}}
		{{- if $.NumFormatIsMillis}}
			time.UnixMilli
		{{- else}}
			time.Unix
		{{- end}}
	{{- else}}
		nil
	{{- end}},
	)
	if err != nil {
		return err
	}
	*t = {{.TyName}}(tt)
	return nil
}

func (t {{.TyName}}) String() string {
    {{if .PreferEpoch -}}
	return strconv.FormatInt(time.Time(t).
		{{- if $.NumFormatIsMillis -}}
			UnixMilli()
		{{- else -}}
			Unix()
		{{- end -}}, 10)
	{{- else -}}
	return time.Time(t).Format(` + "`" + `{{.MarshallingFormat}}` + "`" + `)
	{{- end}}
}
`))

type testTmplParam struct {
	TyName      string
	PackageName string
}

var testTmpl = template.Must(template.New("n").Parse(`
func Fuzz{{.TyName}}(f *testing.F) {
	f.Add(int64(1666282966123))
	f.Fuzz(func(t *testing.T, tValue int64) {
		tt := {{.PackageName}}.{{.TyName}}(time.UnixMilli(tValue))

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
