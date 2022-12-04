package generate

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
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
	if rawTy, ok := fieldTypeTable[prop.Type]; ok {
		return rawTy, GeneratedType{}, nil
	}

	switch prop.Type {
	case mapping.AggregateMetricDouble:
		gen := AggregateMetricDoubleParams(*prop.Param.(*types.AggregateMetricDoubleProperty))
		return gen, GeneratedType{}, nil
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
	case mapping.Date, mapping.DateNanoseconds:
		gen, err := DateFromParam(
			*prop.Param.(*types.DateProperty),
			globalOpt.TypeNameGenerator.Gen(fieldNames),
			opt.PreferredTimeMarshallingFormat,
			opt.PreferTimeEpochMarshalling.True(),
		)
		testDef := DateTest(gen.TyName, "")
		if err != nil {
			return GeneratedType{}, GeneratedType{}, err
		}
		return gen, testDef, nil
	}

	// must not be reached
	return GeneratedType{}, GeneratedType{}, fmt.Errorf("unknown type: %s", prop.Type)
}

var fieldTypeTable = map[mapping.EsType]GeneratedType{
	// FIXME: add special handling for this.
	// Alias type needs knowledge about referenced field...
	// Do nothing here?
	mapping.Alias:           {TyName: "any"},
	mapping.Binary:          {TyName: "[]byte"},
	mapping.Completion:      {TyName: "string"},
	mapping.DenseVector:     {TyName: "[]float64"}, // TODO: read dim and use an array instead of a slice?
	mapping.Flattened:       {TyName: anyMap},
	mapping.Geopoint:        {TyName: estypePrefix + "Geopoint", Imports: estypeImport},
	mapping.Geoshape:        {TyName: estypePrefix + "Geoshape", Imports: estypeImport},
	mapping.IP:              {TyName: "netip.Addr", Imports: []string{`"net/netip"`}},
	mapping.Histogram:       {TyName: anyMap}, // TODO: implement
	mapping.Join:            {TyName: anyMap}, // TODO: implement
	mapping.Percolator:      {TyName: anyMap}, // TODO: implement
	mapping.Point:           {TyName: anyMap}, // TODO: implement
	mapping.RankFeature:     {TyName: "float64"},
	mapping.RankFeatures:    {TyName: float64Map},
	mapping.SearchAsYouType: {TyName: "string"},
	mapping.Shape:           {TyName: estypePrefix + "Geoshape", Imports: estypeImport},
	mapping.TokenCount:      {TyName: "int64"},
	mapping.Version:         {TyName: "string"}, // should this be sem ver package?
	mapping.Keyword:         {TyName: "string"},
	mapping.ConstantKeyword: {TyName: "string"}, // The field can be stored if and only if value is same as specified in param.
	mapping.Wildcard:        {TyName: "string"},
	mapping.Text:            {TyName: "string"},
	// https://www.elastic.co/guide/en/elasticsearch/reference/8.4/number.html
	mapping.Long:         {TyName: "int64"},
	mapping.Integer:      {TyName: "int32"},
	mapping.Short:        {TyName: "int16"},
	mapping.Byte:         {TyName: "int8"}, // The doc says it ranges -128 to 127. It's not the go built-in byte. Rather, it is a typical char type.
	mapping.Double:       {TyName: "float64"},
	mapping.Float:        {TyName: "float32"},
	mapping.HalfFloat:    {TyName: "float32"}, // TODO: use float16 package?
	mapping.UnsignedLong: {TyName: "uint64"},
	mapping.ScaledFloat:  {TyName: "float64"},
	// TODO: implement
	// see https://www.elastic.co/guide/en/elasticsearch/reference/8.4/range.html
	mapping.IntegerRange: {TyName: anyMap},
	mapping.FloatRange:   {TyName: anyMap},
	mapping.LongRange:    {TyName: anyMap},
	mapping.DoubleRange:  {TyName: anyMap},
	mapping.DateRange:    {TyName: anyMap},
	mapping.IpRange:      {TyName: anyMap},
}
