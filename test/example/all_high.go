package example

import (
	"encoding/json"
	"net/netip"
	"time"

	estype "github.com/ngicks/elastic-type/es_type"
	"github.com/ngicks/flextime"
	typeparamcommon "github.com/ngicks/type-param-common"
)

type All struct {
	Agg             *estype.AggregateMetricDouble `json:"agg"`
	Alias           *any                          `json:"alias"`
	Blob            *[]byte                       `json:"blob"`
	Bool            *estype.Boolean               `json:"bool"`
	Byte            *int8                         `json:"byte"`
	Comp            *string                       `json:"comp"`
	ConstantKwd     *string                       `json:"constant_kwd"`
	Date            *AllDate                      `json:"date"`
	DateNano        *AllDateNano                  `json:"dateNano"`
	DateRange       *map[string]interface{}       `json:"date_range"`
	DenseVector     *float64                      `json:"dense_vector"`
	Double          *float64                      `json:"double"`
	DoubleRange     *map[string]interface{}       `json:"double_range"`
	Flattened       *map[string]interface{}       `json:"flattened"`
	Float           *float32                      `json:"float"`
	FloatRange      *map[string]interface{}       `json:"float_range"`
	Geopoint        *estype.Geopoint              `json:"geopoint"`
	Geoshape        *estype.Geoshape              `json:"geoshape"`
	HalfFloat       *float32                      `json:"half_float"`
	Histogram       *map[string]interface{}       `json:"histogram"`
	Integer         *int32                        `json:"integer"`
	IntegerRange    *map[string]interface{}       `json:"integer_range"`
	IpAddr          *netip.Addr                   `json:"ip_addr"`
	IpRange         *map[string]interface{}       `json:"ip_range"`
	Join            *map[string]interface{}       `json:"join"`
	Kwd             *string                       `json:"kwd"`
	Long            *int64                        `json:"long"`
	LongRange       *map[string]interface{}       `json:"long_range"`
	Nested          *AllNested                    `json:"nested"`
	Object          *AllObject                    `json:"object"`
	Point           *map[string]interface{}       `json:"point"`
	Query           *map[string]interface{}       `json:"query"`
	RankFeature     *float64                      `json:"rank_feature"`
	RankFeatures    *map[string]float64           `json:"rank_features"`
	ScaledFloat     *float64                      `json:"scaled_float"`
	SearchAsYouType *string                       `json:"search_as_you_type"`
	Shape           *estype.Geoshape              `json:"shape"`
	Short           *int16                        `json:"short"`
	Text            *string                       `json:"text"`
	TextWTokenCount *string                       `json:"text_w_token_count"`
	UnsignedLong    *uint64                       `json:"unsigned_long"`
	Version         *string                       `json:"version"`
	Wildcard        *string                       `json:"wildcard"`
}

func (t All) ToRaw() AllRaw {
	return AllRaw{
		Agg:          estype.NewFieldSinglePointer(t.Agg, false),
		Alias:        estype.NewFieldSinglePointer(t.Alias, false),
		Blob:         estype.NewFieldSinglePointer(t.Blob, false),
		Bool:         estype.NewFieldSinglePointer(t.Bool, false),
		Byte:         estype.NewFieldSinglePointer(t.Byte, false),
		Comp:         estype.NewFieldSinglePointer(t.Comp, false),
		ConstantKwd:  estype.NewFieldSinglePointer(t.ConstantKwd, false),
		Date:         estype.NewFieldSinglePointer(t.Date, false),
		DateNano:     estype.NewFieldSinglePointer(t.DateNano, false),
		DateRange:    estype.NewFieldSinglePointer(t.DateRange, false),
		DenseVector:  estype.NewFieldSinglePointer(t.DenseVector, false),
		Double:       estype.NewFieldSinglePointer(t.Double, false),
		DoubleRange:  estype.NewFieldSinglePointer(t.DoubleRange, false),
		Flattened:    estype.NewFieldSinglePointer(t.Flattened, false),
		Float:        estype.NewFieldSinglePointer(t.Float, false),
		FloatRange:   estype.NewFieldSinglePointer(t.FloatRange, false),
		Geopoint:     estype.NewFieldSinglePointer(t.Geopoint, false),
		Geoshape:     estype.NewFieldSinglePointer(t.Geoshape, false),
		HalfFloat:    estype.NewFieldSinglePointer(t.HalfFloat, false),
		Histogram:    estype.NewFieldSinglePointer(t.Histogram, false),
		Integer:      estype.NewFieldSinglePointer(t.Integer, false),
		IntegerRange: estype.NewFieldSinglePointer(t.IntegerRange, false),
		IpAddr:       estype.NewFieldSinglePointer(t.IpAddr, false),
		IpRange:      estype.NewFieldSinglePointer(t.IpRange, false),
		Join:         estype.NewFieldSinglePointer(t.Join, false),
		Kwd:          estype.NewFieldSinglePointer(t.Kwd, false),
		Long:         estype.NewFieldSinglePointer(t.Long, false),
		LongRange:    estype.NewFieldSinglePointer(t.LongRange, false),
		Nested: estype.MapField(estype.NewFieldSinglePointer(t.Nested, false), func(v AllNested) AllNestedRaw {
			return v.ToRaw()
		}),
		Object: estype.MapField(estype.NewFieldSinglePointer(t.Object, false), func(v AllObject) AllObjectRaw {
			return v.ToRaw()
		}),
		Point:           estype.NewFieldSinglePointer(t.Point, false),
		Query:           estype.NewFieldSinglePointer(t.Query, false),
		RankFeature:     estype.NewFieldSinglePointer(t.RankFeature, false),
		RankFeatures:    estype.NewFieldSinglePointer(t.RankFeatures, false),
		ScaledFloat:     estype.NewFieldSinglePointer(t.ScaledFloat, false),
		SearchAsYouType: estype.NewFieldSinglePointer(t.SearchAsYouType, false),
		Shape:           estype.NewFieldSinglePointer(t.Shape, false),
		Short:           estype.NewFieldSinglePointer(t.Short, false),
		Text:            estype.NewFieldSinglePointer(t.Text, false),
		TextWTokenCount: estype.NewFieldSinglePointer(t.TextWTokenCount, false),
		UnsignedLong:    estype.NewFieldSinglePointer(t.UnsignedLong, false),
		Version:         estype.NewFieldSinglePointer(t.Version, false),
		Wildcard:        estype.NewFieldSinglePointer(t.Wildcard, false),
	}
}

// AllDate represents elasticsearch date.
type AllDate time.Time

func (t AllDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

var parserAllDate = flextime.NewFlextime(
	typeparamcommon.Must(flextime.NewLayoutSet(`2006-01-02 15:04:05`)).
		AddLayout(typeparamcommon.Must(flextime.NewLayoutSet(`2006-01-02`))),
)

func (t *AllDate) UnmarshalJSON(data []byte) error {
	tt, err := estype.UnmarshalEsTime(
		data,
		parserAllDate.Parse,
		time.UnixMilli,
	)
	if err != nil {
		return err
	}
	*t = AllDate(tt)
	return nil
}

func (t AllDate) String() string {
	return time.Time(t).Format(`2006-01-02 15:04:05`)
}

// AllDateNano represents elasticsearch date.
type AllDateNano time.Time

func (t AllDateNano) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

var parserAllDateNano = flextime.NewFlextime(
	typeparamcommon.Must(flextime.NewLayoutSet(`2006-01-02T15:04:05.999999999Z07:00`)).
		AddLayout(typeparamcommon.Must(flextime.NewLayoutSet(`2006-01-02`))),
)

func (t *AllDateNano) UnmarshalJSON(data []byte) error {
	tt, err := estype.UnmarshalEsTime(
		data,
		parserAllDateNano.Parse,
		func(v int64) time.Time { return time.Unix(v, 0) },
	)
	if err != nil {
		return err
	}
	*t = AllDateNano(tt)
	return nil
}

func (t AllDateNano) String() string {
	return time.Time(t).Format(`2006-01-02T15:04:05.999999999Z07:00`)
}

type AllNested struct {
	Age  *int32   `json:"age"`
	Name *AllName `json:"name"`
}

func (t AllNested) ToRaw() AllNestedRaw {
	return AllNestedRaw{
		Age: estype.NewFieldSinglePointer(t.Age, false),
		Name: estype.MapField(estype.NewFieldSinglePointer(t.Name, false), func(v AllName) AllNameRaw {
			return v.ToRaw()
		}),
	}
}

type AllName struct {
	First *string `json:"first"`
	Last  *string `json:"last"`
}

func (t AllName) ToRaw() AllNameRaw {
	return AllNameRaw{
		First: estype.NewFieldSinglePointer(t.First, false),
		Last:  estype.NewFieldSinglePointer(t.Last, false),
	}
}

type AllObject struct {
	Age  *int32         `json:"age"`
	Name *AllObjectName `json:"name"`
}

func (t AllObject) ToRaw() AllObjectRaw {
	return AllObjectRaw{
		Age: estype.NewFieldSinglePointer(t.Age, false),
		Name: estype.MapField(estype.NewFieldSinglePointer(t.Name, false), func(v AllObjectName) AllObjectNameRaw {
			return v.ToRaw()
		}),
	}
}

type AllObjectName struct {
	First *string `json:"first"`
	Last  *string `json:"last"`
}

func (t AllObjectName) ToRaw() AllObjectNameRaw {
	return AllObjectNameRaw{
		First: estype.NewFieldSinglePointer(t.First, false),
		Last:  estype.NewFieldSinglePointer(t.Last, false),
	}
}
