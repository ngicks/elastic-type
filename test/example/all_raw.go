package example

import (
	"net/netip"

	estype "github.com/ngicks/elastic-type/es_type"
)

type AllRaw struct {
	Agg             estype.Field[estype.AggregateMetricDouble] `json:"agg"`
	Alias           estype.Field[any]                          `json:"alias"`
	Blob            estype.Field[[]byte]                       `json:"blob"`
	Bool            estype.Field[estype.Boolean]               `json:"bool"`
	Byte            estype.Field[int8]                         `json:"byte"`
	Comp            estype.Field[string]                       `json:"comp"`
	ConstantKwd     estype.Field[string]                       `json:"constant_kwd"`
	Date            estype.Field[AllDate]                      `json:"date"`
	DateNano        estype.Field[AllDateNano]                  `json:"dateNano"`
	DateRange       estype.Field[map[string]interface{}]       `json:"date_range"`
	DenseVector     estype.Field[float64]                      `json:"dense_vector"`
	Double          estype.Field[float64]                      `json:"double"`
	DoubleRange     estype.Field[map[string]interface{}]       `json:"double_range"`
	Flattened       estype.Field[map[string]interface{}]       `json:"flattened"`
	Float           estype.Field[float32]                      `json:"float"`
	FloatRange      estype.Field[map[string]interface{}]       `json:"float_range"`
	Geopoint        estype.Field[estype.Geopoint]              `json:"geopoint"`
	Geoshape        estype.Field[estype.Geoshape]              `json:"geoshape"`
	HalfFloat       estype.Field[float32]                      `json:"half_float"`
	Histogram       estype.Field[map[string]interface{}]       `json:"histogram"`
	Integer         estype.Field[int32]                        `json:"integer"`
	IntegerRange    estype.Field[map[string]interface{}]       `json:"integer_range"`
	IpAddr          estype.Field[netip.Addr]                   `json:"ip_addr"`
	IpRange         estype.Field[map[string]interface{}]       `json:"ip_range"`
	Join            estype.Field[map[string]interface{}]       `json:"join"`
	Kwd             estype.Field[string]                       `json:"kwd"`
	Long            estype.Field[int64]                        `json:"long"`
	LongRange       estype.Field[map[string]interface{}]       `json:"long_range"`
	Nested          estype.Field[AllNestedRaw]                 `json:"nested"`
	Object          estype.Field[AllObjectRaw]                 `json:"object"`
	Point           estype.Field[map[string]interface{}]       `json:"point"`
	Query           estype.Field[map[string]interface{}]       `json:"query"`
	RankFeature     estype.Field[float64]                      `json:"rank_feature"`
	RankFeatures    estype.Field[map[string]float64]           `json:"rank_features"`
	ScaledFloat     estype.Field[float64]                      `json:"scaled_float"`
	SearchAsYouType estype.Field[string]                       `json:"search_as_you_type"`
	Shape           estype.Field[estype.Geoshape]              `json:"shape"`
	Short           estype.Field[int16]                        `json:"short"`
	Text            estype.Field[string]                       `json:"text"`
	TextWTokenCount estype.Field[string]                       `json:"text_w_token_count"`
	UnsignedLong    estype.Field[uint64]                       `json:"unsigned_long"`
	Version         estype.Field[string]                       `json:"version"`
	Wildcard        estype.Field[string]                       `json:"wildcard"`
}

func (r AllRaw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t AllRaw) ToPlain() All {
	return All{
		Agg:          t.Agg.Value(),
		Alias:        t.Alias.Value(),
		Blob:         t.Blob.Value(),
		Bool:         t.Bool.Value(),
		Byte:         t.Byte.Value(),
		Comp:         t.Comp.Value(),
		ConstantKwd:  t.ConstantKwd.Value(),
		Date:         t.Date.Value(),
		DateNano:     t.DateNano.Value(),
		DateRange:    t.DateRange.Value(),
		DenseVector:  t.DenseVector.Value(),
		Double:       t.Double.Value(),
		DoubleRange:  t.DoubleRange.Value(),
		Flattened:    t.Flattened.Value(),
		Float:        t.Float.Value(),
		FloatRange:   t.FloatRange.Value(),
		Geopoint:     t.Geopoint.Value(),
		Geoshape:     t.Geoshape.Value(),
		HalfFloat:    t.HalfFloat.Value(),
		Histogram:    t.Histogram.Value(),
		Integer:      t.Integer.Value(),
		IntegerRange: t.IntegerRange.Value(),
		IpAddr:       t.IpAddr.Value(),
		IpRange:      t.IpRange.Value(),
		Join:         t.Join.Value(),
		Kwd:          t.Kwd.Value(),
		Long:         t.Long.Value(),
		LongRange:    t.LongRange.Value(),
		Nested: estype.MapField(t.Nested, func(v AllNestedRaw) AllNested {
			return v.ToPlain()
		}).Value(),
		Object: estype.MapField(t.Object, func(v AllObjectRaw) AllObject {
			return v.ToPlain()
		}).Value(),
		Point:           t.Point.Value(),
		Query:           t.Query.Value(),
		RankFeature:     t.RankFeature.Value(),
		RankFeatures:    t.RankFeatures.Value(),
		ScaledFloat:     t.ScaledFloat.Value(),
		SearchAsYouType: t.SearchAsYouType.Value(),
		Shape:           t.Shape.Value(),
		Short:           t.Short.Value(),
		Text:            t.Text.Value(),
		TextWTokenCount: t.TextWTokenCount.Value(),
		UnsignedLong:    t.UnsignedLong.Value(),
		Version:         t.Version.Value(),
		Wildcard:        t.Wildcard.Value(),
	}
}

type AllNestedRaw struct {
	Age  estype.Field[int32]      `json:"age"`
	Name estype.Field[AllNameRaw] `json:"name"`
}

func (r AllNestedRaw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t AllNestedRaw) ToPlain() AllNested {
	return AllNested{
		Age: t.Age.Value(),
		Name: estype.MapField(t.Name, func(v AllNameRaw) AllName {
			return v.ToPlain()
		}).Value(),
	}
}

type AllNameRaw struct {
	First estype.Field[string] `json:"first"`
	Last  estype.Field[string] `json:"last"`
}

func (r AllNameRaw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t AllNameRaw) ToPlain() AllName {
	return AllName{
		First: t.First.Value(),
		Last:  t.Last.Value(),
	}
}

type AllObjectRaw struct {
	Age  estype.Field[int32]            `json:"age"`
	Name estype.Field[AllObjectNameRaw] `json:"name"`
}

func (r AllObjectRaw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t AllObjectRaw) ToPlain() AllObject {
	return AllObject{
		Age: t.Age.Value(),
		Name: estype.MapField(t.Name, func(v AllObjectNameRaw) AllObjectName {
			return v.ToPlain()
		}).Value(),
	}
}

type AllObjectNameRaw struct {
	First estype.Field[string] `json:"first"`
	Last  estype.Field[string] `json:"last"`
}

func (r AllObjectNameRaw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t AllObjectNameRaw) ToPlain() AllObjectName {
	return AllObjectName{
		First: t.First.Value(),
		Last:  t.Last.Value(),
	}
}
