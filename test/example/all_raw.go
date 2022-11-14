package example

import (
	"net/netip"

	estype "github.com/ngicks/elastic-type/es_type"
)

type AllRaw struct {
	Agg             estype.Field[estype.AggregateMetricDouble] `json:"agg" esjson:"single"`
	Alias           estype.Field[any]                          `json:"alias" esjson:"single"`
	Blob            estype.Field[[]byte]                       `json:"blob" esjson:"single"`
	Bool            estype.Field[estype.Boolean]               `json:"bool" esjson:"single"`
	Byte            estype.Field[int8]                         `json:"byte" esjson:"single"`
	Comp            estype.Field[string]                       `json:"comp" esjson:"single"`
	ConstantKwd     estype.Field[string]                       `json:"constant_kwd" esjson:"single"`
	Date            estype.Field[AllDate]                      `json:"date" esjson:"single"`
	DateNano        estype.Field[AllDateNano]                  `json:"dateNano" esjson:"single"`
	DateRange       estype.Field[map[string]interface{}]       `json:"date_range" esjson:"single"`
	DenseVector     estype.Field[[]float64]                    `json:"dense_vector" esjson:"single"`
	Double          estype.Field[float64]                      `json:"double" esjson:"single"`
	DoubleRange     estype.Field[map[string]interface{}]       `json:"double_range" esjson:"single"`
	Flattened       estype.Field[map[string]interface{}]       `json:"flattened" esjson:"single"`
	Float           estype.Field[float32]                      `json:"float" esjson:"single"`
	FloatRange      estype.Field[map[string]interface{}]       `json:"float_range" esjson:"single"`
	Geopoint        estype.Field[estype.Geopoint]              `json:"geopoint" esjson:"single"`
	Geoshape        estype.Field[estype.Geoshape]              `json:"geoshape" esjson:"single"`
	HalfFloat       estype.Field[float32]                      `json:"half_float" esjson:"single"`
	Histogram       estype.Field[map[string]interface{}]       `json:"histogram" esjson:"single"`
	Integer         estype.Field[int32]                        `json:"integer" esjson:"single"`
	IntegerRange    estype.Field[map[string]interface{}]       `json:"integer_range" esjson:"single"`
	IpAddr          estype.Field[netip.Addr]                   `json:"ip_addr" esjson:"single"`
	IpRange         estype.Field[map[string]interface{}]       `json:"ip_range" esjson:"single"`
	Join            estype.Field[map[string]interface{}]       `json:"join" esjson:"single"`
	Kwd             estype.Field[string]                       `json:"kwd" esjson:"single"`
	Long            estype.Field[int64]                        `json:"long" esjson:"single"`
	LongRange       estype.Field[map[string]interface{}]       `json:"long_range" esjson:"single"`
	Nested          estype.Field[AllNestedRaw]                 `json:"nested" esjson:"single"`
	Object          estype.Field[AllObjectRaw]                 `json:"object" esjson:"single"`
	Point           estype.Field[map[string]interface{}]       `json:"point" esjson:"single"`
	Query           estype.Field[map[string]interface{}]       `json:"query" esjson:"single"`
	RankFeature     estype.Field[float64]                      `json:"rank_feature" esjson:"single"`
	RankFeatures    estype.Field[map[string]float64]           `json:"rank_features" esjson:"single"`
	ScaledFloat     estype.Field[float64]                      `json:"scaled_float" esjson:"single"`
	SearchAsYouType estype.Field[string]                       `json:"search_as_you_type" esjson:"single"`
	Shape           estype.Field[estype.Geoshape]              `json:"shape" esjson:"single"`
	Short           estype.Field[int16]                        `json:"short" esjson:"single"`
	Text            estype.Field[string]                       `json:"text" esjson:"single"`
	TextWTokenCount estype.Field[string]                       `json:"text_w_token_count" esjson:"single"`
	UnsignedLong    estype.Field[uint64]                       `json:"unsigned_long" esjson:"single"`
	Version         estype.Field[string]                       `json:"version" esjson:"single"`
	Wildcard        estype.Field[string]                       `json:"wildcard" esjson:"single"`
}

func (r AllRaw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t AllRaw) ToPlain() All {
	return All{
		Agg:          t.Agg.ValueSingle(),
		Alias:        t.Alias.ValueSingle(),
		Blob:         t.Blob.ValueSingle(),
		Bool:         t.Bool.ValueSingle(),
		Byte:         t.Byte.ValueSingle(),
		Comp:         t.Comp.ValueSingle(),
		ConstantKwd:  t.ConstantKwd.ValueSingle(),
		Date:         t.Date.ValueSingle(),
		DateNano:     t.DateNano.ValueSingle(),
		DateRange:    t.DateRange.ValueSingle(),
		DenseVector:  t.DenseVector.ValueSingle(),
		Double:       t.Double.ValueSingle(),
		DoubleRange:  t.DoubleRange.ValueSingle(),
		Flattened:    t.Flattened.ValueSingle(),
		Float:        t.Float.ValueSingle(),
		FloatRange:   t.FloatRange.ValueSingle(),
		Geopoint:     t.Geopoint.ValueSingle(),
		Geoshape:     t.Geoshape.ValueSingle(),
		HalfFloat:    t.HalfFloat.ValueSingle(),
		Histogram:    t.Histogram.ValueSingle(),
		Integer:      t.Integer.ValueSingle(),
		IntegerRange: t.IntegerRange.ValueSingle(),
		IpAddr:       t.IpAddr.ValueSingle(),
		IpRange:      t.IpRange.ValueSingle(),
		Join:         t.Join.ValueSingle(),
		Kwd:          t.Kwd.ValueSingle(),
		Long:         t.Long.ValueSingle(),
		LongRange:    t.LongRange.ValueSingle(),
		Nested: estype.MapField(t.Nested, func(v AllNestedRaw) AllNested {
			return v.ToPlain()
		}).ValueSingle(),
		Object: estype.MapField(t.Object, func(v AllObjectRaw) AllObject {
			return v.ToPlain()
		}).ValueSingle(),
		Point:           t.Point.ValueSingle(),
		Query:           t.Query.ValueSingle(),
		RankFeature:     t.RankFeature.ValueSingle(),
		RankFeatures:    t.RankFeatures.ValueSingle(),
		ScaledFloat:     t.ScaledFloat.ValueSingle(),
		SearchAsYouType: t.SearchAsYouType.ValueSingle(),
		Shape:           t.Shape.ValueSingle(),
		Short:           t.Short.ValueSingle(),
		Text:            t.Text.ValueSingle(),
		TextWTokenCount: t.TextWTokenCount.ValueSingle(),
		UnsignedLong:    t.UnsignedLong.ValueSingle(),
		Version:         t.Version.ValueSingle(),
		Wildcard:        t.Wildcard.ValueSingle(),
	}
}

type AllNestedRaw struct {
	Age  estype.Field[int32]      `json:"age" esjson:"single"`
	Name estype.Field[AllNameRaw] `json:"name" esjson:"single"`
}

func (r AllNestedRaw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t AllNestedRaw) ToPlain() AllNested {
	return AllNested{
		Age: t.Age.ValueSingle(),
		Name: estype.MapField(t.Name, func(v AllNameRaw) AllName {
			return v.ToPlain()
		}).ValueSingle(),
	}
}

type AllNameRaw struct {
	First estype.Field[string] `json:"first" esjson:"single"`
	Last  estype.Field[string] `json:"last" esjson:"single"`
}

func (r AllNameRaw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t AllNameRaw) ToPlain() AllName {
	return AllName{
		First: t.First.ValueSingle(),
		Last:  t.Last.ValueSingle(),
	}
}

type AllObjectRaw struct {
	Age  estype.Field[int32]            `json:"age" esjson:"single"`
	Name estype.Field[AllObjectNameRaw] `json:"name" esjson:"single"`
}

func (r AllObjectRaw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t AllObjectRaw) ToPlain() AllObject {
	return AllObject{
		Age: t.Age.ValueSingle(),
		Name: estype.MapField(t.Name, func(v AllObjectNameRaw) AllObjectName {
			return v.ToPlain()
		}).ValueSingle(),
	}
}

type AllObjectNameRaw struct {
	First estype.Field[string] `json:"first" esjson:"single"`
	Last  estype.Field[string] `json:"last" esjson:"single"`
}

func (r AllObjectNameRaw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t AllObjectNameRaw) ToPlain() AllObjectName {
	return AllObjectName{
		First: t.First.ValueSingle(),
		Last:  t.Last.ValueSingle(),
	}
}
