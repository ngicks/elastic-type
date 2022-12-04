package generate

import (
	"bytes"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	estype "github.com/ngicks/elastic-type/es_type"
	"github.com/ngicks/elastic-type/mapping"
	"github.com/ngicks/type-param-common/iterator"
)

type tyNameWithOption struct {
	TyName   string
	Option   concreteFieldOption
	HasChild bool
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
	dynamicContext *dynamicmapping.DynamicMapping,
) (highLevelTy, rawTy, testDef []GeneratedType, err error) {
	var subHighLevelTypes, subRawTypes, subTestDefs []GeneratedType
	highLevelFields := map[string]tyNameWithOption{}
	rawFields := map[string]tyNameWithOption{}

	tyName := globalOpt.TypeNameGenerator.Gen(fieldNames)

	// for stable iteration.
	iter := iterator.FromMap(props, func(keys []string) []string {
		sort.Strings(keys)
		return keys
	})

	for _, tuple := range iter.Collect() {
		name, param := tuple.Former, tuple.Latter
		fieldOption := opts[name]
		overlaidOption := globalOpt.Overlay(param, fieldOption)

		if param.IsObjectLike() {
			var subHighLevelTy, subRawTy, subTestDef []GeneratedType
			var err error

			if param.IsObject() {
				subHighLevelTy, subRawTy, subTestDef, err = Object(
					*param.Param.(*mapping.ObjectProperty),
					globalOpt,
					fieldOption.ChildOption,
					append(fieldNames, name),
					dynamicContext,
				)
			} else {
				subHighLevelTy, subRawTy, subTestDef, err = Nested(
					*param.Param.(*mapping.NestedProperty),
					globalOpt,
					fieldOption.ChildOption,
					append(fieldNames, name),
					dynamicContext,
				)
			}
			if err != nil {
				return nil, nil, nil, err
			}

			subHighLevelTypes = append(subHighLevelTypes, subHighLevelTy...)
			subRawTypes = append(subRawTypes, subRawTy...)
			subTestDefs = append(subTestDefs, subTestDef...)

			subHighLevelTy[0].Option = overlaidOption
			subRawTy[0].Option = overlaidOption

			highLevelFields[name] = tyNameWithOption{
				TyName:   subHighLevelTy[0].TyName,
				Option:   fieldOptToConcrete(overlaidOption),
				HasChild: true,
			}
			rawFields[name] = tyNameWithOption{
				TyName:   subRawTy[0].TyName,
				Option:   fieldOptToConcrete(overlaidOption),
				HasChild: true,
			}

		} else {
			gen, testDef, err := Field(param, append(fieldNames, name), globalOpt, fieldOption)
			gen.Option = overlaidOption

			if err != nil {
				return nil, nil, nil, err
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
			subRawTypes = append(subRawTypes, GeneratedType{Imports: gen.Imports})
			subTestDefs = append(subTestDefs, testDef)
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
		TyName:  tyName,
		TyDef:   buf.String(),
		Imports: estypeImport,
	}

	return append([]GeneratedType{thisType}, subHighLevelTypes...),
		append([]GeneratedType{thisTypeRaw}, subRawTypes...),
		subTestDefs,
		nil
}

func Object(
	p mapping.ObjectProperty,
	globalOpt GlobalOption,
	opts MapOption,
	fieldNames []string,
	dynamicContext dynamicmapping.DynamicMapping,
) (highLevelTy, rawTy, testDef []GeneratedType, err error) {
	newDynamic := mapping.OverlayDynamic(dynamicContext, p.Dynamic)

	// only when dynamic == true. ignore other cases (e.g. "runtime", false)
	if mapping.DynamicIsTrue(newDynamic) {
		// What should we do it when Dynamic is "runtime"?
		// TODO: research what will happen then.
		tyName := capitalize(globalOpt.TypeNameGenerator.Gen(fieldNames))

		params := struct {
			TyName    string
			RawTyName string
		}{
			TyName:    tyName,
			RawTyName: tyName + "Raw",
		}

		highDef := bytes.NewBuffer(make([]byte, 0))
		rawDef := bytes.NewBuffer(make([]byte, 0))
		var err error
		err = objectHighMapTemplate.Execute(highDef, params)
		if err != nil {
			panic(err)
		}
		err = objectRawMapTemplate.Execute(rawDef, params)
		if err != nil {
			panic(err)
		}

		return []GeneratedType{
				{
					TyName: tyName,
					TyDef:  highDef.String(),
				},
			},
			[]GeneratedType{
				{
					TyName: tyName + "Raw",
					TyDef:  rawDef.String(),
				},
			},
			[]GeneratedType{},
			nil
	}
	return object(*p.Properties, globalOpt, opts, fieldNames, newDynamic)
}

var objectHighMapTemplate = template.Must(template.New("objectHighMapTemplate").Parse(`
type {{.TyName}} map[string][]any

func (t {{.TyName}}) ToRaw() {{.RawTyName}} {
	out := {{.RawTyName}}{}
	for k, v := range t {
		out[k] = estype.NewFieldSlice(v, false)
	}
	return out
}`))

var objectRawMapTemplate = template.Must(template.New("objectRawMapTemplate").Parse(`
type {{.RawTyName}} map[string]estype.Field[any]

func (t {{.RawTyName}}) ToPlain() {{.TyName}} {
	out := {{.TyName}}{}
	for k, v := range t {
		out[k] = v.ValueZero()
	}
	return out
}`))

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

func (r {{.TyName}}Raw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t {{.TyName}}Raw) ToPlain() {{.TyName}} {
	return {{.TyName}}{
{{range $propName, $typeNameOpt := .HighLevelFields}}` +
	// field name: value methods
	`    {{toPascalCase $propName}}: 
	{{- if $typeNameOpt.HasChild -}}
		estype.MapField(t.{{toPascalCase $propName}}, func(v {{with $rawField := index $.RawFields $propName }}{{$rawField.TyName}}{{end}}) {{$typeNameOpt.TyName}} {
			return v.ToPlain()
		})
	{{- else -}}
		t.{{toPascalCase $propName}}
	{{- end -}}.` +
	`{{- if $typeNameOpt.Option.IsSingle -}}
		{{- if $typeNameOpt.Option.IsRequired -}}ValueSingleZero()
		{{- else -}}ValueSingle()
		{{- end -}}
	{{- else -}}
		{{- if $typeNameOpt.Option.IsRequired -}}ValueZero()
		{{- else -}}Value()
		{{- end -}}
	{{- end }},` +
	`
{{end}}
	}
}
`))

var objectTemplate = template.Must(template.New("objectTemplate").Funcs(funcMap).Parse(`
type {{.TyName}} struct {
{{range $propName, $typeNameOpt := .HighLevelFields}}` +
	`    {{toPascalCase $propName}}   {{ if not $typeNameOpt.Option.IsRequired -}}*{{ end }}{{ if not $typeNameOpt.Option.IsSingle }}[]{{ end }}{{$typeNameOpt.TyName}}   ` + "`" + `json:"{{$propName}}"` + "`" +
	`
{{end }}}

func (t {{.TyName}}) ToRaw() {{.TyName}}Raw {
	return {{.TyName}}Raw{
{{range $propName, $typeNameOpt := .RawFields}}` +
	// field name: value methods
	`    {{toPascalCase $propName}}: 
	{{- if $typeNameOpt.HasChild -}}
		estype.MapField(
	{{- end -}}	` +
	`{{- if $typeNameOpt.Option.IsSingle -}}
		{{- if $typeNameOpt.Option.IsRequired -}}estype.NewFieldSingleValue(t.{{toPascalCase $propName}}) {{/* T */}}
		{{- else -}}estype.NewFieldSinglePointer(t.{{toPascalCase $propName}}, false) {{/* *T */}}
		{{- end -}}
	{{- else -}}
		{{- if $typeNameOpt.Option.IsRequired -}}estype.NewFieldSlice(t.{{toPascalCase $propName}}, false) {{/* []T */}}
		{{- else -}}estype.NewField(t.{{toPascalCase $propName}}) {{/* *[]T */}}
		{{- end -}}
	{{- end }}` + `			
	{{- if $typeNameOpt.HasChild -}}, func(v {{with $rawField := index $.HighLevelFields $propName }}{{$rawField.TyName}}{{end}}) {{$typeNameOpt.TyName}} {
			return v.ToRaw()
		})
	{{- end}},` +
	`
{{end}}		
	}
}
`))
