package mapping

import (
	"bytes"
	"encoding/json"
)

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
//
// seeing these documents reveals that mappings key is actually the object field mapping type.
//   - https://www.elastic.co/guide/en/elasticsearch/reference/8.4/dynamic.html
//   - https://www.elastic.co/guide/en/elasticsearch/reference/8.4/enabled.html
//   - https://www.elastic.co/guide/en/elasticsearch/reference/8.4/subobjects.html
//   - https://www.elastic.co/guide/en/elasticsearch/reference/8.4/properties.html
type Mappings = ObjectParams

type Properties map[string]Property

func (p *Properties) FillType() {
	for _, v := range *p {
		if filler, ok := v.Param.(FillTyper); ok {
			filler.FillType()
		}
	}
}

type Property struct {
	Type  EsType
	Param any
}

func (p Property) IsObject() bool {
	return p.Type == Object || p.Type == ""
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

type Dynamic = json.RawMessage

func IsValidDynamic(d Dynamic) bool {
	if d == nil {
		// defaults to true.
		return true
	}

	for _, dynamic := range validDynamic {
		if dynamic == nil {
			// bytes.Equal is just a string(a) == string(a).
			// We now enforce that nil is not []byte(``)
			continue
		}
		if bytes.Equal(d, dynamic) {
			return true
		}
	}
	return false
}

func OverlayDynamic(left, right Dynamic) Dynamic {
	if len(right) != 0 && IsValidDynamic(right) {
		return right
	}
	if IsValidDynamic(left) {
		return left
	}
	return Empty
}

func DynamicIsTrue(d Dynamic) bool {
	return string([]byte(d)) == string(TrueBool) || string([]byte(d)) == string(TrueStr)
}

var (
	Empty     Dynamic = nil
	TrueBool  Dynamic = []byte(`true`)
	FalseBool Dynamic = []byte(`false`)
	TrueStr   Dynamic = []byte(`"true"`)
	FalseStr  Dynamic = []byte(`"false"`)
	Runtime   Dynamic = []byte(`"runtime"`)
	Strict    Dynamic = []byte(`"script"`)
)

var validDynamic = []Dynamic{
	Empty,
	TrueBool,
	FalseBool,
	TrueStr,
	FalseStr,
	Runtime,
	Strict,
}
