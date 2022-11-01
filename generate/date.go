package generate

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	builtinformat "github.com/ngicks/elastic-type/es_type/builtin_format"
	"github.com/ngicks/elastic-type/mapping"
	"github.com/ngicks/flextime"
	"github.com/ngicks/type-param-common/set"
	"github.com/ngicks/type-param-common/slice"
)

// DateFromParam generates type from given parameters.
//
// If tyName is empty, generated type name is EsDateTime.
// If marshallingFormat is empty, longest format will be used as marshaller.
// If preferEpochMarshalling is true, generated type marshals into epoch millis or second. string otherwize.
//
// Currently prop is used only for Formats. Future update may use other fields.
func DateFromParam(
	prop mapping.DateParams,
	tyName string,
	marshallingFormat string,
	preferEpochMarhsalling bool,
) (generatedTyName string, tyDef string, imports []string, err error) {
	var formats string
	if prop.Format == nil {
		// must be sync with
		return estypePrefix + "StrictDateOptionalTimeEpochMillis", "", estypeImport, nil
	} else {
		formats = *prop.Format
	}
	layouts, hasNumFormat, isMillis, err := ParseFormatsString(formats)
	if err != nil {
		return "", "", nil, err
	}

	if marshallingFormat == "" {
		marshallingFormat = layouts.Layout()[0]
	} else {
		converted, err := flextime.ReplaceTimeToken(marshallingFormat)
		if err != nil {
			return "", "", nil, err
		}
		if !slice.Has(layouts.Layout(), converted) {
			return "", "", nil, fmt.Errorf(
				"prefered formated %s is not one of formats %+v",
				marshallingFormat,
				layouts.Layout(),
			)
		}
		marshallingFormat = converted
	}

	// TODO: check if format is only one and is one of elasticsearch built-in formats.
	// If so, don't generate type. Use estype.<Type> instead.
	generatedTyName, tyDef, imports = DateUnchecked(DateGenerationParam{
		TyName:            typeName(tyName),
		StrFormats:        layouts.Layout(),
		MarshallingFormat: marshallingFormat,
		HasNumFormat:      hasNumFormat,
		NumFormatIsMillis: isMillis,
		PreferEpoch:       preferEpochMarhsalling,
	})
	return
}

func typeName(n string) string {
	if n == "" {
		return "EsDateTime"
	} else {
		return n
	}
}

func DateUnchecked(params DateGenerationParam) (tyName string, tyDef string, imports []string) {
	buf := bytes.NewBuffer(make([]byte, 0))
	err := dateTypeTmpl.Execute(buf, params)
	if err != nil {
		panic(err)
	}

	tyName = params.TyName
	tyDef = buf.String()
	imports = generateIpmorts(params.PreferEpoch)

	return tyName, tyDef, imports
}

// DateTest generates test for a type which is result of generate.Date().
// tyNamePrefix must be same of passed to generate.Date().
// pkgName must be package name containing that generated type.
func DateTest(tyNamePrefix string, pkgName string) (tyName string, tyDef string) {
	buf := bytes.NewBuffer(make([]byte, 0))

	_ = dateTestTmpl.Execute(buf, DateTestTmplParam{
		TyName:      typeName(tyNamePrefix),
		PackageName: pkgName,
	})

	return tyName, buf.String()
}

func ParseFormatsString(formats string) (layouts *flextime.LayoutSet, hasNumFormat, isMillis bool, err error) {
	formatsSl := strings.Split(formats, "||")
	return ParseFormats(formatsSl)
}

func ParseFormats(formats []string) (layouts *flextime.LayoutSet, hasNumFormat, isMillis bool, err error) {
	strFormats, hasNumFormat, isMillis, _ := excludeNumFormats(formats)

	first, rest := strFormats[0], strFormats[1:]

	layoutSet, err := flextime.NewLayoutSet(first)
	if err != nil {
		return nil, false, false, err
	}
	for _, format := range rest {
		additive, err := flextime.NewLayoutSet(format)
		if err != nil {
			return nil, false, false, err
		}
		layoutSet = layoutSet.AddLayout(additive)
	}
	return layoutSet, hasNumFormat, isMillis, nil
}

func excludeNumFormats(formats []string) (strFormats []string, hasNumFormat, isMillis, hasDupe bool) {
	strFormats = make([]string, 0)
	formatSet := set.New[string]()
	for _, format := range formats {
		v, ok := builtinformat.ParsingLayout[format]

		if ok {
			formatSet.Add(v)
			strFormats = append(strFormats, v)
		} else if format == builtinformat.EpochMillis || format == builtinformat.EpochSecond {
			if hasNumFormat {
				hasDupe = true
			}

			hasNumFormat = true
			if format == builtinformat.EpochMillis {
				isMillis = true
			}
		} else {
			formatSet.Add(format)
			strFormats = append(strFormats, format)
		}
	}

	if len(strFormats) != formatSet.Len() {
		hasDupe = true
	}

	return formatSet.Values().Collect(), hasNumFormat, isMillis, hasDupe
}

func generateIpmorts(preferEpoch bool) []string {
	imports := make([]string, 0)
	if preferEpoch {
		imports = append(imports, `"strconv"`)
	} else {
		imports = append(imports, `"encoding/json"`)
	}
	imports = append(imports, []string{
		`"time"`,
		`estype "github.com/ngicks/elastic-type/es_type"`,
		`"github.com/ngicks/flextime"`,
		`typeparamcommon "github.com/ngicks/type-param-common"`,
	}...)
	return imports
}

type DateGenerationParam struct {
	TyName            string   // Name of type.
	StrFormats        []string // Unmarshalling formats excluding epoch_millis or epoch_second. format must be consist of go lang time tokens.
	MarshallingFormat string   // format used in String and MarshalJSON (used only when PreferEpoch is false).
	HasNumFormat      bool     // has epoch_millis or epoch_second
	NumFormatIsMillis bool     // format is epoch_millis
	PreferEpoch       bool     // marshal into number json value.
}

var dateTypeTmpl = template.Must(template.New("v").Parse(`
// {{.TyName}} represents elasticsearch date.
type {{.TyName}} time.Time

func (t {{.TyName}}) MarshalJSON() ([]byte, error) {
{{- if .PreferEpoch}}
	return []byte(t.String()), nil
{{else}}
	return json.Marshal(t.String())
{{- end -}}

}


var parser{{.TyName}} = flextime.NewFlextime(
{{range $index, $format := .StrFormats}}
	{{- if eq $index 0 }}	typeparamcommon.Must(flextime.NewLayoutSet(` + "`" + `{{$format}}` + "`" + `))
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
		func(v int64) time.Time { return time.Unix(v, 0) }
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

type DateTestTmplParam struct {
	TyName      string
	PackageName string
}

var dateTestTmpl = template.Must(template.New("n").Parse(`
	f.Add(int64(1666282966123), int64(218964089023))
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
