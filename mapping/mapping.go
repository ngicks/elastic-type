package mapping

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
)

// IndexSettings is main body for [Create index API.](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/indices-create-index.html)
type IndexSettings struct {
	types.IndexState
	Mappings *TypeMapping `json:"mappings,omitempty"`
}

// TypeMapping is main body for [updating mapping](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/explicit-mapping.html#add-field-mapping),
// or a part of [Create index API](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/indices-create-index.html).
type TypeMapping struct {
	types.TypeMapping
	Properties *Properties `json:"properties,omitempty"`
}

// MappingSettings is response body of GET /<index_name>/_mapping.
//
// Map may only contain <index_name> as a key, and the contained IndexSettings may only have Mappings.
type MappingSettings map[string]IndexSettings

type Properties map[string]Property

type Property struct {
	Type  EsType
	Param any
}

func (p Property) IsObject() bool {
	return p.Type == Object || p.Type == ""
}

func (p Property) IsObjectLike() bool {
	return p.IsObject() || p.Type == Nested
}

func (p Property) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Param)
}

func (p *Property) UnmarshalJSON(data []byte) error {
	// TODO: use efficient way to retrieve type key from data.
	type Ty struct {
		Type EsType `json:"type,omitempty"`
	}
	ty := Ty{}

	var err error
	err = json.Unmarshal(data, &ty)
	if err != nil {
		return err
	}

	if ty.Type != "" {
		p.Type = ty.Type
	}

	switch ty.Type {
	default:
		p.Type = Object

		var o ObjectProperty
		p.Param = &o
	case AggregateMetricDouble:
		var o types.AggregateMetricDoubleProperty
		p.Param = &o
	case Alias:
		var o types.FieldAliasProperty
		p.Param = &o
	case Binary:
		var o types.BinaryProperty
		p.Param = &o
	case Boolean:
		var o types.BooleanProperty
		p.Param = &o
	case Completion:
		var o types.CompletionProperty
		p.Param = &o
	case Date, DateNanoseconds:
		var o types.DateProperty
		p.Param = &o
	case DenseVector:
		var o types.DenseVectorProperty
		p.Param = &o
	case Flattened:
		var o types.FlattenedProperty
		p.Param = &o
	case Geopoint:
		var o types.GeoPointProperty
		p.Param = &o
	case Geoshape:
		var o types.GeoShapeProperty
		p.Param = &o
	case Histogram:
		var o types.HistogramProperty
		p.Param = &o
	case IP:
		var o types.IpProperty
		p.Param = &o
	case Join:
		var o types.JoinProperty
		p.Param = &o
	case Nested:
		var o NestedProperty
		p.Param = &o
	case Percolator:
		var o types.PercolatorProperty
		p.Param = &o
	case Point:
		var o types.PointProperty
		p.Param = &o
	case RankFeature, RankFeatures:
		var o types.RankFeatureProperty
		p.Param = &o
	case SearchAsYouType:
		var o types.SearchAsYouTypeProperty
		p.Param = &o
	case Shape:
		var o types.ShapeProperty
		p.Param = &o
	case TokenCount:
		var o types.TokenCountProperty
		p.Param = &o
	case Version:
		var o types.VersionProperty
		p.Param = &o
	case Keyword:
		var o types.KeywordProperty
		p.Param = &o
	case ConstantKeyword:
		var o types.ConstantKeywordProperty
		p.Param = &o
	case Wildcard:
		var o types.WildcardProperty
		p.Param = &o
	case Text:
		var o types.TextProperty
		p.Param = &o
	case Long, Integer, Short, Byte, Double, Float, HalfFloat, UnsignedLong:
		var o types.NumberPropertyBase
		p.Param = &o
	case ScaledFloat:
		var o types.ScaledFloatNumberProperty
		p.Param = &o
	case IntegerRange, FloatRange, LongRange, DoubleRange:
		var o types.RangePropertyBase
		p.Param = &o
	case DateRange:
		var o types.DateRangeProperty
		p.Param = &o
	case IpRange:
		var o types.IpRangeProperty
		p.Param = &o
	}

	err = json.Unmarshal(data, &p.Param)
	if err != nil {
		return err
	}
	return nil
}

var possibleDynamic = [...]dynamicmapping.DynamicMapping{
	dynamicmapping.Strict,
	dynamicmapping.Runtime,
	dynamicmapping.True,
	dynamicmapping.False,
}

var emptyDynamic = dynamicmapping.DynamicMapping{}

func IsEmptyDynamic(d dynamicmapping.DynamicMapping) bool {
	return emptyDynamic == d
}

func IsValidDynamic(d dynamicmapping.DynamicMapping) bool {
	for _, v := range possibleDynamic {
		if v == d {
			return true
		}
	}

	return false
}

func OverlayDynamic(left, right dynamicmapping.DynamicMapping) dynamicmapping.DynamicMapping {
	if IsValidDynamic(right) {
		return right
	}
	if IsValidDynamic(left) {
		return left
	}
	return emptyDynamic
}
