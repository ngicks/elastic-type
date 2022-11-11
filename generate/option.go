package generate

import (
	"fmt"
	"strconv"

	"github.com/ngicks/elastic-type/mapping"
	"github.com/ngicks/type-param-common/set"
)

type optStr string

const (
	Inherit optStr = ""
	None    optStr = ""
	True    optStr = "true"
	False   optStr = "false"
)

func (s optStr) True() bool {
	return s == True
}

func (s optStr) False() bool {
	return s == False
}

func (s optStr) Empty() bool {
	return s == None
}

func (s optStr) Overlay(other optStr) optStr {
	if !other.Empty() {
		return other
	} else {
		return s
	}
}

type TypeNameGenerationRule func(fieldNames []string) string

func FieldName() TypeNameGenerationRule {
	return func(fieldNames []string) string {
		return fieldNames[len(fieldNames)-1]
	}
}

type Haser interface {
	Has(string) bool
}

type TypeNameFallBackRule func(tyName string, fieldNames []string, usedTypeName Haser) string

func UseOneUpperFieldName(shouldPanicOnOverlap bool) TypeNameFallBackRule {
	return func(tyName string, fieldNames []string, usedTypeName Haser) string {
		if len(fieldNames) < 2 {
			panic("broken invariants: fieldNames must be len(c) >= 2")
		}
		tyName = capitalize(fieldNames[len(fieldNames)-2]) + capitalize(tyName)

		if usedTypeName.Has(tyName) {
			if shouldPanicOnOverlap {
				panic(fmt.Errorf("overlapping name: %s", tyName))
			}

			old := tyName
			var i int64 = 1
			for {
				tyName = old + strconv.FormatInt(i, 10)
				if !usedTypeName.Has(tyName) {
					break
				}
				i++
			}
		}

		return tyName
	}
}

type TypeNamePostProcessRule func(string) string

func PascalCaseUnderscoreHyphen() TypeNamePostProcessRule {
	return func(s string) string {
		return toPascalCaseDelimiter(s)
	}
}

type TypeNameGenerator struct {
	Generate     TypeNameGenerationRule
	FallBack     TypeNameFallBackRule
	PostProcess  TypeNamePostProcessRule
	usedTypeName *set.Set[string]
}

func (g *TypeNameGenerator) lazyInit() {
	if g.Generate == nil {
		g.Generate = FieldName()
	}
	if g.FallBack == nil {
		g.FallBack = UseOneUpperFieldName(false)
	}
	if g.PostProcess == nil {
		g.PostProcess = PascalCaseUnderscoreHyphen()
	}
	if g.usedTypeName == nil {
		g.usedTypeName = set.New[string]()
	}
}

func (g *TypeNameGenerator) Gen(fieldNames []string) string {
	g.lazyInit()

	tyName := g.Generate(fieldNames)

	if g.usedTypeName.Has(tyName) {
		tyName = g.FallBack(tyName, fieldNames, g.usedTypeName)
	}

	g.usedTypeName.Add(tyName)
	return g.PostProcess(tyName)
}

type GlobalOption struct {
	IsRequired                 optStr            // prefer fields to be unmarshalled into non-pointer type T, instead of *T.
	IsSingle                   optStr            // prefer fields to be unmarshalled into single value T, instead of []T.
	PreferStringBoolean        optStr            // prefer Boolean types to marshal into "true" / "false".
	PreferTimeEpochMarshalling optStr            // prefer Date types to marshal into epoch millis or epoch second.
	TypeOption                 TypeOption        // Default options for the type.
	TypeNameGenerator          TypeNameGenerator // Defaults to FieldName().
}

// Overlay overlays options.
//
// The priorities are global-option < type-option < field-option.
func (g GlobalOption) Overlay(prop mapping.Property, fieldOpt FieldOption) FieldOption {
	if g.TypeOption == nil {
		g.TypeOption = TypeOption{}
	}

	return FieldOption{
		IsRequired: g.IsRequired.
			Overlay(g.TypeOption[prop.Type].IsRequired).
			Overlay(fieldOpt.IsRequired),
		IsSingle: g.IsSingle.
			Overlay(g.TypeOption[prop.Type].IsSingle).
			Overlay(fieldOpt.IsSingle),
		PreferStringBoolean: g.PreferStringBoolean.Overlay(
			fieldOpt.PreferStringBoolean,
		),
		PreferredTimeMarshallingFormat: fieldOpt.PreferredTimeMarshallingFormat,
		PreferTimeEpochMarshalling: g.PreferTimeEpochMarshalling.Overlay(
			fieldOpt.PreferTimeEpochMarshalling,
		),
	}
}

type OptionForType struct {
	IsRequired optStr // prefer fields to be unmarshalled into non-pointer type T, instead of *T.
	IsSingle   optStr // prefer fields to be unmarshalled into single value T, instead of []T.
}

type TypeOption map[mapping.EsType]OptionForType

// GetDefaultTypeOption returns an opinionated defaults for TypeOption.
func GetDefaultTypeOption() TypeOption {
	// clone default to avoid mutation.
	to := TypeOption{}
	for k, v := range defaultTypeOption {
		to[k] = v
	}
	return to
}

var defaultTypeOption = TypeOption{
	mapping.AggregateMetricDouble: OptionForType{
		IsSingle: True,
	},
	mapping.Alias: OptionForType{},
	mapping.Binary: OptionForType{
		IsSingle: True,
	},
	mapping.Boolean: OptionForType{
		IsSingle: True,
	},
	mapping.Completion: OptionForType{
		IsSingle: False,
	},
	mapping.Date: OptionForType{
		IsSingle: True,
	},
	mapping.DateNanoseconds: OptionForType{
		IsSingle: True,
	},
	mapping.DenseVector: OptionForType{},
	mapping.Flattened:   OptionForType{}, // TODO: research
	mapping.Geopoint: OptionForType{
		IsSingle: False,
	},
	mapping.Geoshape: OptionForType{
		IsSingle: False,
	},
	mapping.Histogram: OptionForType{}, // TODO: research
	mapping.IP:        OptionForType{},
	mapping.Join:      OptionForType{},
	mapping.Nested: OptionForType{
		IsRequired: True,
		IsSingle:   False,
	},
	mapping.Object: OptionForType{
		IsRequired: True,
		IsSingle:   False,
	},
	mapping.Percolator: OptionForType{},
	mapping.Point: OptionForType{
		IsSingle: False,
	},
	mapping.RankFeature:     OptionForType{},
	mapping.RankFeatures:    OptionForType{},
	mapping.SearchAsYouType: OptionForType{},
	mapping.Shape:           OptionForType{},
	mapping.TokenCount:      OptionForType{},
	mapping.Version:         OptionForType{},
	mapping.Keyword:         OptionForType{},
	mapping.ConstantKeyword: OptionForType{},
	mapping.Wildcard:        OptionForType{},
	mapping.Text:            OptionForType{},
	mapping.Long: OptionForType{
		IsSingle: True,
	},
	mapping.Integer: OptionForType{
		IsSingle: True,
	},
	mapping.Short: OptionForType{
		IsSingle: True,
	},
	mapping.Byte: OptionForType{
		IsSingle: True,
	},
	mapping.Double: OptionForType{
		IsSingle: True,
	},
	mapping.Float: OptionForType{
		IsSingle: True,
	},
	mapping.HalfFloat: OptionForType{
		IsSingle: True,
	},
	mapping.ScaledFloat: OptionForType{
		IsSingle: True,
	},
	mapping.UnsignedLong: OptionForType{
		IsSingle: True,
	},
	mapping.IntegerRange: OptionForType{},
	mapping.FloatRange:   OptionForType{},
	mapping.LongRange:    OptionForType{},
	mapping.DoubleRange:  OptionForType{},
	mapping.DateRange:    OptionForType{},
	mapping.IpRange:      OptionForType{},
}

// Options for Object, Nested type or a root element of an Elasticsearch mapping.
type MapOption map[string]FieldOption

type FieldOption struct {
	IsRequired                     optStr
	IsSingle                       optStr
	PreferStringBoolean            optStr
	PreferredTimeMarshallingFormat string // no inheritance for this field.
	PreferTimeEpochMarshalling     optStr
	ChildOption                    MapOption
}
