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
	Agg             *[]estype.AggregateMetricDouble `json:"agg"`
	Alias           *[]any                          `json:"alias"`
	Blob            *[][]byte                       `json:"blob"`
	Bool            *[]estype.Boolean               `json:"bool"`
	Byte            *[]int8                         `json:"byte"`
	Comp            *[]string                       `json:"comp"`
	ConstantKwd     *[]string                       `json:"constant_kwd"`
	Date            *[]AllDate                      `json:"date"`
	DateNano        *[]AllDateNano                  `json:"dateNano"`
	DateRange       *[]map[string]interface{}       `json:"date_range"`
	DenseVector     *[]float64                      `json:"dense_vector"`
	Double          *[]float64                      `json:"double"`
	DoubleRange     *[]map[string]interface{}       `json:"double_range"`
	Flattened       *[]map[string]interface{}       `json:"flattened"`
	Float           *[]float32                      `json:"float"`
	FloatRange      *[]map[string]interface{}       `json:"float_range"`
	Geopoint        *[]estype.Geopoint              `json:"geopoint"`
	Geoshape        *[]estype.Geoshape              `json:"geoshape"`
	HalfFloat       *[]float32                      `json:"half_float"`
	Histogram       *[]map[string]interface{}       `json:"histogram"`
	Integer         *[]int32                        `json:"integer"`
	IntegerRange    *[]map[string]interface{}       `json:"integer_range"`
	IpAddr          *[]netip.Addr                   `json:"ip_addr"`
	IpRange         *[]map[string]interface{}       `json:"ip_range"`
	Join            *[]map[string]interface{}       `json:"join"`
	Kwd             *[]string                       `json:"kwd"`
	Long            *[]int64                        `json:"long"`
	LongRange       *[]map[string]interface{}       `json:"long_range"`
	Nested          *[]AllNested                    `json:"nested"`
	Object          *[]AllObject                    `json:"object"`
	Point           *[]map[string]interface{}       `json:"point"`
	Query           *[]map[string]interface{}       `json:"query"`
	RankFeature     *[]float64                      `json:"rank_feature"`
	RankFeatures    *[]map[string]float64           `json:"rank_features"`
	ScaledFloat     *[]float64                      `json:"scaled_float"`
	SearchAsYouType *[]string                       `json:"search_as_you_type"`
	Shape           *[]estype.Geoshape              `json:"shape"`
	Short           *[]int16                        `json:"short"`
	Text            *[]string                       `json:"text"`
	TextWTokenCount *[]string                       `json:"text_w_token_count"`
	UnsignedLong    *[]uint64                       `json:"unsigned_long"`
	Version         *[]string                       `json:"version"`
	Wildcard        *[]string                       `json:"wildcard"`
}

func (t All) ToRaw() AllRaw {
	return AllRaw{
		Agg:          estype.NewField(t.Agg, false),
		Alias:        estype.NewField(t.Alias, false),
		Blob:         estype.NewField(t.Blob, false),
		Bool:         estype.NewField(t.Bool, false),
		Byte:         estype.NewField(t.Byte, false),
		Comp:         estype.NewField(t.Comp, false),
		ConstantKwd:  estype.NewField(t.ConstantKwd, false),
		Date:         estype.NewField(t.Date, false),
		DateNano:     estype.NewField(t.DateNano, false),
		DateRange:    estype.NewField(t.DateRange, false),
		DenseVector:  estype.NewField(t.DenseVector, false),
		Double:       estype.NewField(t.Double, false),
		DoubleRange:  estype.NewField(t.DoubleRange, false),
		Flattened:    estype.NewField(t.Flattened, false),
		Float:        estype.NewField(t.Float, false),
		FloatRange:   estype.NewField(t.FloatRange, false),
		Geopoint:     estype.NewField(t.Geopoint, false),
		Geoshape:     estype.NewField(t.Geoshape, false),
		HalfFloat:    estype.NewField(t.HalfFloat, false),
		Histogram:    estype.NewField(t.Histogram, false),
		Integer:      estype.NewField(t.Integer, false),
		IntegerRange: estype.NewField(t.IntegerRange, false),
		IpAddr:       estype.NewField(t.IpAddr, false),
		IpRange:      estype.NewField(t.IpRange, false),
		Join:         estype.NewField(t.Join, false),
		Kwd:          estype.NewField(t.Kwd, false),
		Long:         estype.NewField(t.Long, false),
		LongRange:    estype.NewField(t.LongRange, false),
		Nested: estype.MapField(estype.NewField(t.Nested, false), func(v AllNested) AllNestedRaw {
			return v.ToRaw()
		}),
		Object: estype.MapField(estype.NewField(t.Object, false), func(v AllObject) AllObjectRaw {
			return v.ToRaw()
		}),
		Point:           estype.NewField(t.Point, false),
		Query:           estype.NewField(t.Query, false),
		RankFeature:     estype.NewField(t.RankFeature, false),
		RankFeatures:    estype.NewField(t.RankFeatures, false),
		ScaledFloat:     estype.NewField(t.ScaledFloat, false),
		SearchAsYouType: estype.NewField(t.SearchAsYouType, false),
		Shape:           estype.NewField(t.Shape, false),
		Short:           estype.NewField(t.Short, false),
		Text:            estype.NewField(t.Text, false),
		TextWTokenCount: estype.NewField(t.TextWTokenCount, false),
		UnsignedLong:    estype.NewField(t.UnsignedLong, false),
		Version:         estype.NewField(t.Version, false),
		Wildcard:        estype.NewField(t.Wildcard, false),
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
	Age  *[]int32   `json:"age"`
	Name *[]AllName `json:"name"`
}

func (t AllNested) ToRaw() AllNestedRaw {
	return AllNestedRaw{
		Age: estype.NewField(t.Age, false),
		Name: estype.MapField(estype.NewField(t.Name, false), func(v AllName) AllNameRaw {
			return v.ToRaw()
		}),
	}
}

type AllName struct {
	First *[]string `json:"first"`
	Last  *[]string `json:"last"`
}

func (t AllName) ToRaw() AllNameRaw {
	return AllNameRaw{
		First: estype.NewField(t.First, false),
		Last:  estype.NewField(t.Last, false),
	}
}

type AllObject struct {
	Age  *[]int32         `json:"age"`
	Name *[]AllObjectName `json:"name"`
}

func (t AllObject) ToRaw() AllObjectRaw {
	return AllObjectRaw{
		Age: estype.NewField(t.Age, false),
		Name: estype.MapField(estype.NewField(t.Name, false), func(v AllObjectName) AllObjectNameRaw {
			return v.ToRaw()
		}),
	}
}

type AllObjectName struct {
	First *[]string `json:"first"`
	Last  *[]string `json:"last"`
}

func (t AllObjectName) ToRaw() AllObjectNameRaw {
	return AllObjectNameRaw{
		First: estype.NewField(t.First, false),
		Last:  estype.NewField(t.Last, false),
	}
}
