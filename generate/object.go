package generate

import (
	"bytes"
	"regexp"
	"strings"
	"text/template"

	"github.com/ngicks/elastic-type/mapping"
)

func object(props mapping.Properties, opts Options, currentPointer []string) (highLevelTy, rawTy []GeneratedType, err error) {
	var subHighLevelTypes, subRawTypes []GeneratedType
	highLevelFields := map[string]string{}
	rawFields := map[string]string{}

	for name, param := range props {
		option := opts[name]

		if param.IsObject() || param.Type == mapping.Nested {
			var subHighLevelTy, subRawTy []GeneratedType
			var err error

			if param.IsObject() {
				subHighLevelTy, subRawTy, err = Object(
					*param.Param.(*mapping.ObjectParams),
					option.ChildOption,
					append(currentPointer, name),
				)
			} else {
				subHighLevelTy, subRawTy, err = Nested(
					*param.Param.(*mapping.NestedParams),
					option.ChildOption,
					append(currentPointer, name),
				)
			}
			if err != nil {
				return nil, nil, err
			}

			subHighLevelTypes = append(subHighLevelTypes, subHighLevelTy...)
			subRawTypes = append(subRawTypes, subRawTy...)

			highLevelFields[name] = subHighLevelTy[0].TyName
			rawFields[name] = subRawTy[0].TyName
		} else {
			gen, err := Field(param, currentPointer, option)
			if err != nil {
				return nil, nil, err
			}

			highLevelFields[name] = gen.TyName
			rawFields[name] = gen.TyName

			subHighLevelTypes = append(subHighLevelTypes, gen)
			subRawTypes = append(subRawTypes, gen)
		}
	}

	fieldName := currentPointer[len(currentPointer)-1]
	tyName := toPascalCaseDelimiter(fieldName)

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

func Object(p mapping.ObjectParams, opts Options, currentPointer []string) (highLevelTy, rawTy []GeneratedType, err error) {
	return object(*p.Properties, opts, currentPointer)
}

var caseDelimiter = regexp.MustCompile("[_-]")

func capitalize(v string) string {
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
	HighLevelFields map[string]string
	RawFields       map[string]string
}

var funcMap = template.FuncMap{
	"toPascalCase": toPascalCaseDelimiter,
}

var objectRawTemplate = template.Must(template.New("objectRawTemplate").Funcs(funcMap).Parse(`
type {{.TyName}}Raw struct {
{{range $propName, $typeName := .RawFields}}
	{{toPascalCase $propName}}    estype.Field[{{$typeName}}]   ` + "`" + `json:"{{$propName}}"` + "`" + `
{{end}}
}
`))

var objectTemplate = template.Must(template.New("objectTemplate").Funcs(funcMap).Parse(`
type {{.TyName}} struct {
{{range $propName, $typeName := .HighLevelFields}}
	{{toPascalCase $propName}}    {{$typeName}}   ` + "`" + `json:"{{$propName}}"` + "`" + `
{{end}}
}
`))
