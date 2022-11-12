package generate

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"text/template"

	estype "github.com/ngicks/elastic-type/es_type"
	"github.com/ngicks/elastic-type/mapping"
)

type tyNameWithOption struct {
	TyName string
	Option concreteFieldOption
}

type concreteFieldOption struct {
	IsRequired                     bool
	IsSingle                       bool
	PreferStringBoolean            bool
	PreferredTimeMarshallingFormat string
	PreferTimeEpochMarshalling     bool
}

func fieldOptToConcrete(f FieldOption) concreteFieldOption {
	return concreteFieldOption{
		IsRequired:                     f.IsRequired.True(),
		IsSingle:                       f.IsSingle.True(),
		PreferStringBoolean:            f.PreferStringBoolean.True(),
		PreferredTimeMarshallingFormat: f.PreferredTimeMarshallingFormat,
		PreferTimeEpochMarshalling:     f.PreferTimeEpochMarshalling.True(),
	}
}

func object(
	props mapping.Properties,
	globalOpt GlobalOption,
	opts MapOption,
	fieldNames []string,
) (highLevelTy, rawTy []GeneratedType, err error) {
	var subHighLevelTypes, subRawTypes []GeneratedType
	highLevelFields := map[string]tyNameWithOption{}
	rawFields := map[string]tyNameWithOption{}

	tyName := globalOpt.TypeNameGenerator.Gen(fieldNames)

	for name, param := range props {
		fieldOption := opts[name]
		overlaidOption := globalOpt.Overlay(param, fieldOption)

		if param.IsObject() || param.Type == mapping.Nested {
			var subHighLevelTy, subRawTy []GeneratedType
			var err error

			if param.IsObject() {
				subHighLevelTy, subRawTy, err = Object(
					*param.Param.(*mapping.ObjectParams),
					globalOpt,
					fieldOption.ChildOption,
					append(fieldNames, name),
				)
			} else {
				subHighLevelTy, subRawTy, err = Nested(
					*param.Param.(*mapping.NestedParams),
					globalOpt,
					fieldOption.ChildOption,
					append(fieldNames, name),
				)
			}
			if err != nil {
				return nil, nil, err
			}

			subHighLevelTypes = append(subHighLevelTypes, subHighLevelTy...)
			subRawTypes = append(subRawTypes, subRawTy...)

			subHighLevelTy[0].Option = overlaidOption
			subRawTy[0].Option = overlaidOption

			highLevelFields[name] = tyNameWithOption{
				TyName: subHighLevelTy[0].TyName,
				Option: fieldOptToConcrete(overlaidOption),
			}
			rawFields[name] = tyNameWithOption{
				TyName: subRawTy[0].TyName,
				Option: fieldOptToConcrete(overlaidOption),
			}

		} else {
			gen, err := Field(param, append(fieldNames, name), globalOpt, fieldOption)
			gen.Option = overlaidOption

			if err != nil {
				return nil, nil, err
			}

			highLevelFields[name] = tyNameWithOption{
				TyName: gen.TyName,
				Option: fieldOptToConcrete(overlaidOption),
			}
			rawFields[name] = tyNameWithOption{
				TyName: gen.TyName,
				Option: fieldOptToConcrete(overlaidOption),
			}

			subHighLevelTypes = append(subHighLevelTypes, gen)
		}
	}

	buf := bytes.NewBuffer(make([]byte, 0))
	param := objectTemplateParam{
		TyName:          tyName,
		HighLevelFields: highLevelFields,
		RawFields:       rawFields,
	}
	err = objectRawTemplate.Execute(buf, param)
	if err != nil {
		panic(err)
	}

	thisTypeRaw := GeneratedType{
		TyName:  tyName + "Raw",
		TyDef:   buf.String(),
		Imports: estypeImport,
	}

	buf.Reset()
	err = objectTemplate.Execute(buf, param)
	if err != nil {
		panic(err)
	}

	thisType := GeneratedType{
		TyName: tyName,
		TyDef:  buf.String(),
	}

	return append([]GeneratedType{thisType}, subHighLevelTypes...),
		append([]GeneratedType{thisTypeRaw}, subRawTypes...),
		nil
}

func Object(
	p mapping.ObjectParams,
	globalOpt GlobalOption,
	opts MapOption,
	fieldNames []string,
) (highLevelTy, rawTy []GeneratedType, err error) {
	// Ignore dynamic inheritance.
	if p.Dynamic != nil && (*p.Dynamic == "true" || *p.Dynamic == true) {
		// What should we do it when Dynamic is "runtime"?
		// TODO: research what will happen then.
		tyName := capitalize(globalOpt.TypeNameGenerator.Gen(fieldNames))
		return []GeneratedType{
				{
					TyName: tyName,
					TyDef:  fmt.Sprintf("type %s map[string]any", tyName),
				},
			}, []GeneratedType{
				{
					TyName: tyName + "Raw",
					TyDef:  fmt.Sprintf("type %s map[string]any", tyName+"Raw"),
				},
			}, nil
	}
	return object(*p.Properties, globalOpt, opts, fieldNames)
}

var caseDelimiter = regexp.MustCompile("[_-]")

func capitalize(v string) string {
	if length := len(v); length == 0 {
		return v
	} else if length == 1 {
		return strings.ToUpper(v)
	}

	head, rest := v[:1], v[1:]
	return strings.ToUpper(head) + rest
}

func toPascalCase(v []string) string {
	capitalized := make([]string, len(v))
	copy(capitalized, v)
	for i := 0; i < len(v); i++ {
		capitalized[i] = capitalize(capitalized[i])
	}

	return strings.Join(capitalized, "")
}

func toPascalCaseDelimiter(v string) string {
	return toPascalCase(caseDelimiter.Split(v, -1))
}

type objectTemplateParam struct {
	TyName          string
	HighLevelFields map[string]tyNameWithOption
	RawFields       map[string]tyNameWithOption
}

var funcMap = template.FuncMap{
	"toPascalCase": toPascalCaseDelimiter,
}

var objectRawTemplate = template.Must(template.New("objectRawTemplate").Funcs(funcMap).Parse(`
type {{.TyName}}Raw struct {
{{range $propName, $typeNameOpt := .RawFields}}` +
	// field name - field type
	`    {{toPascalCase $propName}}    estype.Field[{{$typeNameOpt.TyName}}]   ` +
	// struct tag
	"`" + `json:"{{$propName}}"  {{- if $typeNameOpt.Option.IsSingle }}` +
	` ` + estype.StructTag + `:"` + estype.TagSingle + `"{{ end -}}` + "`" +
	`
{{end}}}
`))

var objectTemplate = template.Must(template.New("objectTemplate").Funcs(funcMap).Parse(`
type {{.TyName}} struct {
{{range $propName, $typeNameOpt := .HighLevelFields}}` +
	`    {{toPascalCase $propName}}   {{ if not $typeNameOpt.Option.IsRequired -}}*{{ end }}{{ if not $typeNameOpt.Option.IsSingle }}[]{{ end }}{{$typeNameOpt.TyName}}   ` + "`" + `json:"{{$propName}}"` + "`" +
	`
{{end }}}
`))
