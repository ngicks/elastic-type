package test_test

import (
	"net/netip"
	"testing"
	"time"

	"github.com/go-spatial/geom"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	estype "github.com/ngicks/elastic-type/es_type"
	"github.com/ngicks/elastic-type/test"
	"github.com/ngicks/elastic-type/test/example"
	"github.com/ngicks/gommon/pkg/randstr"
	tpc "github.com/ngicks/type-param-common"
	"github.com/stretchr/testify/require"
)

func TestGenerate_all(t *testing.T) {
	require := require.New(t)

	skipIfEsNotReachable(t, *ELASTICSEARCH_URL, false)
	helper := must(createRandomIndex[example.AllRaw](client, test.AllMappings))
	defer helper.Delete()

	randomStr := func() string {
		str, err := randstr.New(randstr.Hex()).String()
		if err != nil {
			panic(err)
		}
		return str
	}
	nowNano := time.Now()
	nowSec := time.Unix(nowNano.Unix(), 0)

	t.Run("store and retrieve to exact same data", func(t *testing.T) {
		allPlain := example.All{
			Agg:         tpc.Escape(estype.AggregateMetricDouble{}),
			Blob:        tpc.Escape([]byte(`foobarbaz`)),
			Bool:        tpc.Escape(estype.Boolean(true)),
			Byte:        tpc.Escape(int8(12)),
			Comp:        tpc.Escape(randomStr()),
			ConstantKwd: tpc.Escape("debug"),
			Date:        tpc.Escape(example.AllDate(nowSec)),
			DateNano:    tpc.Escape(example.AllDateNano(nowNano)),
			DateRange: tpc.Escape(map[string]interface{}{
				"gte": float64(12345),
				"lte": float64(12350),
			}),
			DenseVector: tpc.Escape([]float64{16, 15, 14}),
			Double:      tpc.Escape(float64(68)),
			DoubleRange: tpc.Escape(map[string]interface{}{
				"gte": 10.1,
				"lt":  20.1,
			}),
			Flattened: tpc.Escape(map[string]interface{}{
				"priority": "urgent",
				"release":  []any{"v1.2.5", "v1.3.0"},
				"timestamp": map[string]any{
					"created": float64(1541458026),
					"closed":  float64(1541457010),
				},
			}),
			Float: tpc.Escape(float32(357.3209)),
			FloatRange: tpc.Escape(map[string]interface{}{
				"gte": 10.1,
				"lt":  20.1,
			}),
			Geopoint: tpc.Escape(estype.Geopoint{
				Lat: 41.12,
				Lon: -71.34,
			}),
			Geoshape: tpc.Escape(estype.Geoshape{
				Geometry: geom.Point{-77.03653, 38.897676},
			}),
			HalfFloat: tpc.Escape(float32(2131.57)),
			Histogram: tpc.Escape(map[string]interface{}{
				"values": []any{0.1, 0.2, 0.3, 0.4, 0.5},
				"counts": []any{float64(3), float64(7), float64(23), float64(12), float64(6)},
			}),
			Integer: tpc.Escape(int32(60)),
			IntegerRange: tpc.Escape(map[string]interface{}{
				"gte": float64(10),
				"lt":  float64(20),
			}),
			IpAddr: tpc.Escape(netip.MustParseAddr("192.168.0.1")),
			IpRange: tpc.Escape(map[string]interface{}{
				"gte": "192.168.0.2",
				"lt":  "192.168.0.240",
			}),
			Join: tpc.Escape(map[string]interface{}{
				"name": "question",
			}),
			Kwd:  tpc.Escape("naaaaaaaaaaaaaah"),
			Long: tpc.Escape(int64(210389467827)),
			LongRange: tpc.Escape(map[string]interface{}{
				"gte": float64(10),
				"lt":  float64(20),
			}),
			Nested: tpc.Escape(example.AllNested{
				Age: tpc.Escape(int32(123)),
				Name: &example.AllName{
					First: tpc.Escape("john"),
					Last:  tpc.Escape("doe"),
				},
			}),
			Object: tpc.Escape(example.AllObject{
				Age: tpc.Escape(int32(123)),
				Name: &example.AllObjectName{
					First: tpc.Escape("john"),
					Last:  tpc.Escape("doe"),
				},
			}),
			Point: tpc.Escape(map[string]interface{}{
				"type":        "Point",
				"coordinates": []any{-71.34, 41.12},
			}),
			Query: tpc.Escape(map[string]interface{}{
				"match": map[string]any{
					"kwd": "value",
				},
			}),
			RankFeature: tpc.Escape(float64(124.6)),
			RankFeatures: tpc.Escape(map[string]float64{
				"politics":  float64(20),
				"economics": 50.8,
			}),
			ScaledFloat:     tpc.Escape(float64(12315.4798)),
			SearchAsYouType: tpc.Escape("quick brown fox jump lazy dog"),
			Shape: tpc.Escape(estype.Geoshape{
				Geometry: geom.Point{-77.03653, 38.897676},
			}),
			Short:           tpc.Escape(int16(2109)),
			Text:            tpc.Escape("fox fox fox"),
			TextWTokenCount: tpc.Escape("1208956i;lzcxjo"),
			UnsignedLong:    tpc.Escape(uint64(2109381027538706718)),
			Version:         tpc.Escape("1.2.7"),
			Wildcard:        tpc.Escape("8lnmkvlouiejhr02983"),
		}
		id, err := helper.PostDoc(allPlain.ToRaw())
		require.NoError(err)

		fetchedDoc, err := helper.GetDoc(id)
		require.NoErrorf(err, "%s", err)

		fetchedPlain := fetchedDoc.Source_.ToPlain()

		diff := cmp.Diff(
			fetchedPlain,
			allPlain,
			cmpopts.IgnoreFields(allPlain, "Date", "DateNano", "IpAddr"),
		)
		if diff != "" {
			require.Failf("diff = %s", diff)
		}

		require.True(time.Time(*allPlain.Date).Equal(time.Time(*fetchedPlain.Date)))
		require.True(time.Time(*allPlain.DateNano).Equal(time.Time(*fetchedPlain.DateNano)))
		require.Equal(allPlain.IpAddr.String(), fetchedPlain.IpAddr.String())
	})
}
