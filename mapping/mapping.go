package mapping

import "encoding/json"

// IndexSettings is main body for [Create index API.](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/indices-create-index.html)
type IndexSettings struct {
	// Aliases is irrevant for this package's goal.
	// But allowing it to be stored in here is maybe useful for some user.
	Aliases any `json:"aliases,omitempty"`
	// Settings is irrevant for this package's goal.
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

		// Here you can assign a reference of concrete type, such as ObjectParams, to p.Param (seemingly it works as intended.)
		// And you can reduce these repetitive json.Unmarshal call.
		// But we want to keep it non-pointer. So let it be for now.
		var o ObjectParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case AggregateMetricDouble:
		var o AggregateMetricDoubleParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case Alias:
		var o AliasParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case Binary:
		var o BinaryParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case Boolean:
		var o BooleanParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case Completion:
		var o CompletionParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case Date, DateNanoseconds:
		var o DateParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case DenseVector:
		var o DenseVectorParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case Flattened:
		var o FlattenedParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case Geopoint:
		var o GeopointParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case Geoshape:
		var o GeoshapeParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case Histogram:
		var o HistogramParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case IP:
		var o IPParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case Join:
		var o JoinParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case Nested:
		var o NestedParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case Percolator:
		var o PercolatorParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case Point:
		var o PointParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case Range:
		var o RangeParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case RankFeature, RankFeatures:
		var o RankFeatureParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case SearchAsYouType:
		var o SearchAsYouTypeParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case Shape:
		var o ShapeParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case TokenCount:
		var o TokenCountParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case Version:
		var o VersionParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case Keyword:
		var o KeywordParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case ConstantKeyword:
		var o ConstantKeywordParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case Wildcard:
		var o WildcardParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case Text:
		var o TextParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case Long, Integer, Short, Byte, Double, Float, HalfFloat, UnsignedLong:
		var o NumericParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case ScaledFloat:
		var o ScaledFloatParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	case IntegerRange, FloatRange, LongRange, DoubleRange, DateRange, IpRange:
		var o RangeParams
		err = json.Unmarshal(data, &o)
		p.Param = o
	}

	p.Param.(FillTyper).FillType()

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
