package generate

import (
	"fmt"

	"github.com/ngicks/elastic-type/mapping"
)

// prefix including dot
const estypePrefix = "estype."

const anyMap = "map[string]interface{}"

var estypeImport = []string{`estype "github.com/ngicks/elastic-type/es_type"`}

type GenerationOption struct {
	PreferStringBoolean bool
	PreferedDateFormat  string
	PreferEpochDate     bool
}

// GenerateFieldType generate a type for a corresponding field.
// tyDef will be empty string if tyName is one of go built-in types (e.g. string, int, uint64).
// imports will be nil if tyDef if empty or it does not depends on other packages.
func GenerateFieldType(prop mapping.Property, tyNamePrefix string, opt GenerationOption) (tyName string, tyDef string, imports []string) {
	switch prop.Type {
	default:
		if prop.Type != "" && prop.Type != mapping.Object {
			panic(fmt.Sprintf("unknown: %s", prop.Type))
		}
		// TODO: do recursive type generation.
		tyName = anyMap
		imports = nil
	case mapping.AggregateMetricDouble:
		tyName, tyDef = GenerateAggregateMetricDoubleParams(
			prop.Param.(mapping.AggregateMetricDoubleParams),
			tyNamePrefix,
		)
		imports = nil
	case mapping.Alias:
		// FIXME: add special handling for this.
		// Alias type needs knowledge about referenced field...
		// Do nothing here?
	case mapping.Binary:
		tyName = estypePrefix + "Binary"
		imports = estypeImport
	case mapping.Boolean:
		if opt.PreferStringBoolean {
			tyName = estypePrefix + "BooleanStr"
		} else {
			tyName = estypePrefix + "Boolean"
		}
		imports = estypeImport
	case mapping.Completion:
		tyName = "string"
	case mapping.Date, mapping.DateNanoseconds:
		// validity must be checked in caller. err is ignored here.
		tyName, tyDef, imports, _ = DateFromParam(
			prop.Param.(mapping.DateParams),
			tyNamePrefix+"DateTime",
			opt.PreferedDateFormat,
			opt.PreferEpochDate,
		)
	case mapping.DenseVector:
		tyName = "float64"
	case mapping.Flattened:
		tyName = anyMap
	case mapping.Geopoint:
		tyName = estypePrefix + "Geopoint"
		imports = estypeImport
	case mapping.Geoshape:
		tyName = estypePrefix + "Geoshape"
		imports = estypeImport
	case mapping.IP:
		tyName = "netip.Addr"
		imports = []string{"net/netip"}
	case mapping.Histogram, mapping.Join, mapping.Nested, mapping.Percolator, mapping.Point:
		// TODO: implement
		tyName = anyMap
	case mapping.RankFeature:
		tyName = "float64"
	case mapping.RankFeatures:
		tyName = "map[string]float64"
	case mapping.SearchAsYouType:
		tyName = "string"
	case mapping.Shape:
		tyName = estypePrefix + "Geoshape"
		imports = estypeImport
	case mapping.TokenCount:
		tyName = "int64"
	case mapping.Version:
		// should this be sem ver package?
		tyName = "string"
	case mapping.Keyword:
		tyName = "string"
	case mapping.ConstantKeyword:
		// This field should not be stored?
		tyName = "string"
	case mapping.Wildcard:
		tyName = "string"
	case mapping.Text:
		tyName = "string"
	case mapping.Long, mapping.Integer, mapping.Short, mapping.Byte, mapping.Double, mapping.Float, mapping.HalfFloat, mapping.UnsignedLong:
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
			tyName = "float16"
		case mapping.UnsignedLong:
			tyName = "uint64"
		}
	case mapping.ScaledFloat:
		tyName = "float64"
	case mapping.IntegerRange, mapping.FloatRange, mapping.LongRange, mapping.DoubleRange, mapping.DateRange, mapping.IpRange:
		// TODO: implement
		// see https://www.elastic.co/guide/en/elasticsearch/reference/8.4/range.html
		tyName = anyMap
	}

	return
}
