package mapping

import "encoding/json"

// IndexSettings is main body for [Create index API.](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/indices-create-index.html)
type IndexSettings struct {
	// Aliases is irrelevant for this package's goal.
	// But allowing it to be stored in here is maybe useful for some user.
	Aliases any `json:"aliases,omitempty"`
	// Settings is irrelevant for this package's goal.
	// But allowing it to be stored in here is maybe useful for some user.
	Settings any       `json:"settings,omitempty"`
	Mappings *Mappings `json:"mappings,omitempty"`
}

// MappingSettings is response body of GET /<index_name>/_mapping.
//
// Map may only contain <index_name> as a key, and the contained IndexSettings may only have Mappings.
type MappingSettings map[string]IndexSettings

// Mappings is main body for [updating mapping](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/explicit-mapping.html#add-field-mapping),
// or a part of [Create index API](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/indices-create-index.html).
type Mappings struct {
	Properties Properties `json:"properties"`
}

type Properties map[string]Property

func (p *Properties) FillType() {
	for _, v := range *p {
		if filler, ok := v.Param.(FillTyper); ok {
			filler.FillType()
		}
	}
}

type Property struct {
	Type  esType
	Param any
}

func (p Property) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Param)
}

func (p *Property) UnmarshalJSON(data []byte) error {
	// TODO: use efficient way to retrieve type key from data.
	type Ty struct {
		Type esType `json:"type,omitempty"`
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

		var o ObjectParams
		p.Param = &o
	case AggregateMetricDouble:
		var o AggregateMetricDoubleParams
		p.Param = &o
	case Alias:
		var o AliasParams
		p.Param = &o
	case Binary:
		var o BinaryParams
		p.Param = &o
	case Boolean:
		var o BooleanParams
		p.Param = &o
	case Completion:
		var o CompletionParams
		p.Param = &o
	case Date, DateNanoseconds:
		var o DateParams
		p.Param = &o
	case DenseVector:
		var o DenseVectorParams
		p.Param = &o
	case Flattened:
		var o FlattenedParams
		p.Param = &o
	case Geopoint:
		var o GeopointParams
		p.Param = &o
	case Geoshape:
		var o GeoshapeParams
		p.Param = &o
	case Histogram:
		var o HistogramParams
		p.Param = &o
	case IP:
		var o IPParams
		p.Param = &o
	case Join:
		var o JoinParams
		p.Param = &o
	case Nested:
		var o NestedParams
		p.Param = &o
	case Percolator:
		var o PercolatorParams
		p.Param = &o
	case Point:
		var o PointParams
		p.Param = &o
	case RankFeature, RankFeatures:
		var o RankFeatureParams
		p.Param = &o
	case SearchAsYouType:
		var o SearchAsYouTypeParams
		p.Param = &o
	case Shape:
		var o ShapeParams
		p.Param = &o
	case TokenCount:
		var o TokenCountParams
		p.Param = &o
	case Version:
		var o VersionParams
		p.Param = &o
	case Keyword:
		var o KeywordParams
		p.Param = &o
	case ConstantKeyword:
		var o ConstantKeywordParams
		p.Param = &o
	case Wildcard:
		var o WildcardParams
		p.Param = &o
	case Text:
		var o TextParams
		p.Param = &o
	case Long, Integer, Short, Byte, Double, Float, HalfFloat, UnsignedLong:
		var o NumericParams
		p.Param = &o
	case ScaledFloat:
		var o ScaledFloatParams
		p.Param = &o
	case IntegerRange, FloatRange, LongRange, DoubleRange, DateRange, IpRange:
		var o RangeParams
		p.Param = &o
	}

	err = json.Unmarshal(data, &p.Param)
	if err != nil {
		return err
	}
	return nil
}

type FillTyper interface {
	// FillType fills Type field if it is zero value.
	FillType()
}

type onScriptError string

const (
	Continue onScriptError = "continue"
	Fail     onScriptError = "fail"
)
