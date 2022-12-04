package generate

import (
	"fmt"

	"github.com/ngicks/elastic-type/mapping"
	"github.com/ngicks/type-param-common/slice"
)

// prefix including dot
const estypePrefix = "estype."

const (
	anyMap     = "map[string]interface{}"
	float64Map = "map[string]float64"
)

var estypeImport = []string{`estype "github.com/ngicks/elastic-type/es_type"`}

// Field generates a type for input property.
// Input prop must be one that can not be nested (other than Object or Nested types).
func Field(
	prop mapping.Property,
	fieldNames slice.Deque[string],
	globalOpt GlobalOption,
	opt FieldOption,
) (rawTy, testDef GeneratedType, err error) {
	switch prop.Type {
	case mapping.AggregateMetricDouble:
		gen := AggregateMetricDoubleParams(*prop.Param.(*mapping.AggregateMetricDoubleParams))
		return gen, GeneratedType{}, nil
	case mapping.Alias:
		// FIXME: add special handling for this.
		// Alias type needs knowledge about referenced field...
		// Do nothing here?
		return GeneratedType{
				TyName: "any",
			},
			GeneratedType{},
			nil
	case mapping.Binary:
		return GeneratedType{
				TyName: "[]byte",
			},
			GeneratedType{},
			nil
	case mapping.Boolean:
		var tyName string
		if opt.PreferStringBoolean.True() {
			tyName = estypePrefix + "BooleanStr"
		} else {
			tyName = estypePrefix + "Boolean"
		}

		return GeneratedType{
				TyName:  tyName,
				Imports: estypeImport,
			},
			GeneratedType{},
			nil
	case mapping.Completion:
		return GeneratedType{
				TyName: "string",
			},
			GeneratedType{},
			nil
	case mapping.Date, mapping.DateNanoseconds:
		gen, err := DateFromParam(
			*prop.Param.(*mapping.DateParams),
			globalOpt.TypeNameGenerator.Gen(fieldNames),
			opt.PreferredTimeMarshallingFormat,
			opt.PreferTimeEpochMarshalling.True(),
		)
		testDef := DateTest(gen.TyName, "")
		if err != nil {
			return GeneratedType{}, GeneratedType{}, err
		}
		return gen, testDef, nil
	case mapping.DenseVector:
		return GeneratedType{
				// TODO: read dim and use an array instead of a slice?
				TyName: "[]float64",
			},
			GeneratedType{},
			nil
	case mapping.Flattened:
		return GeneratedType{
				TyName: anyMap,
			},
			GeneratedType{},
			nil
	case mapping.Geopoint:
		return GeneratedType{
				TyName:  estypePrefix + "Geopoint",
				Imports: estypeImport,
			},
			GeneratedType{},
			nil
	case mapping.Geoshape:
		return GeneratedType{
				TyName:  estypePrefix + "Geoshape",
				Imports: estypeImport,
			},
			GeneratedType{},
			nil
	case mapping.IP:
		return GeneratedType{
				TyName:  "netip.Addr",
				Imports: []string{`"net/netip"`},
			},
			GeneratedType{},
			nil
	case mapping.Histogram, mapping.Join, mapping.Percolator, mapping.Point:
		// TODO: implement
		return GeneratedType{
				TyName: anyMap,
			},
			GeneratedType{},
			nil
	case mapping.RankFeature:
		return GeneratedType{
				TyName: "float64",
			},
			GeneratedType{},
			nil
	case mapping.RankFeatures:
		return GeneratedType{
				TyName: float64Map,
			},
			GeneratedType{},
			nil
	case mapping.SearchAsYouType:
		return GeneratedType{
				TyName: "string",
			},
			GeneratedType{},
			nil
	case mapping.Shape:
		return GeneratedType{
				TyName:  estypePrefix + "Geoshape",
				Imports: estypeImport,
			},
			GeneratedType{},
			nil
	case mapping.TokenCount:
		return GeneratedType{
				TyName: "int64",
			},
			GeneratedType{},
			nil
	case mapping.Version:
		// should this be sem ver package?
		return GeneratedType{
				TyName: "string",
			},
			GeneratedType{},
			nil
	case mapping.Keyword:
		return GeneratedType{
				TyName: "string",
			},
			GeneratedType{},
			nil
	case mapping.ConstantKeyword:
		// The field can be stored if and only if value is same as specified in param.
		return GeneratedType{
				TyName: "string",
			},
			GeneratedType{},
			nil
	case mapping.Wildcard:
		return GeneratedType{
				TyName: "string",
			},
			GeneratedType{},
			nil
	case mapping.Text:
		return GeneratedType{
				TyName: "string",
			},
			GeneratedType{},
			nil
	case mapping.Long, mapping.Integer, mapping.Short, mapping.Byte, mapping.Double, mapping.Float, mapping.HalfFloat, mapping.UnsignedLong:
		var tyName string
		// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/number.html
		switch prop.Type {
		case mapping.Long:
			tyName = "int64"
		case mapping.Integer:
			tyName = "int32"
		case mapping.Short:
			tyName = "int16"
		case mapping.Byte:
			// The doc says it ranges -128 to 127. It's not the go built-in byte. Rather, it is a typical char type.
			tyName = "int8"
		case mapping.Double:
			tyName = "float64"
		case mapping.Float:
			tyName = "float32"
		case mapping.HalfFloat:
			// there is not float16 type in built-in types.
			// TODO: use float16 package?
			tyName = "float32"
		case mapping.UnsignedLong:
			tyName = "uint64"
		}

		return GeneratedType{
				TyName: tyName,
			},
			GeneratedType{},
			nil
	case mapping.ScaledFloat:
		return GeneratedType{
				TyName: "float64",
			},
			GeneratedType{},
			nil
	case mapping.IntegerRange, mapping.FloatRange, mapping.LongRange, mapping.DoubleRange, mapping.DateRange, mapping.IpRange:
		// TODO: implement
		// see https://www.elastic.co/guide/en/elasticsearch/reference/8.4/range.html
		return GeneratedType{
				TyName: anyMap,
			},
			GeneratedType{},
			nil
	}

	// must not be reached
	return GeneratedType{}, GeneratedType{}, fmt.Errorf("unknown type: %s", prop.Type)
}
