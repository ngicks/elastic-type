package estype_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/go-spatial/geom"
	"github.com/google/go-cmp/cmp"
	estype "github.com/ngicks/elastic-type/es_type"
	"github.com/stretchr/testify/require"
)

func TestGeoshape(t *testing.T) {
	require := require.New(t)

	points := [][]byte{
		[]byte(`{"type" : "Point", "coordinates" : [-77.03653, 38.897676]}`),
		[]byte(`"POINT (-77.03653 38.897676)"`),
	}

	for _, bin := range points {
		var g estype.Geoshape

		err := json.Unmarshal(bin, &g)
		require.NoError(err)

		point := g.Geometry.(geom.Point)
		require.InDelta(-77.03653, point.X(), 0.001)
		require.InDelta(38.897676, point.Y(), 0.001)

		bin, err := json.Marshal(g)
		require.NoError(err)

		// it does marshal into geojson.
		require.True(bytes.HasPrefix(bin, []byte(`{`)))
		require.True(bytes.HasSuffix(bin, []byte(`}`)))

		var g2 estype.Geoshape

		err = json.Unmarshal(bin, &g2)
		require.NoError(err)

		point = g2.Geometry.(geom.Point)
		require.InDelta(-77.03653, point.X(), 0.001)
		require.InDelta(38.897676, point.Y(), 0.001)
	}

	lines := [][]byte{
		[]byte(`{
			"type" : "LineString",
			"coordinates" : [[-77.03653, 38.897676], [-77.009051, 38.889939]]
		}`),
		[]byte(`"LINESTRING (-77.03653 38.897676, -77.009051 38.889939)"`),
	}

	for _, bin := range lines {
		var g estype.Geoshape

		err := json.Unmarshal(bin, &g)
		require.NoError(err)

		line := g.Geometry.(geom.LineString)

		expected := [][2]float64{{-77.03653, 38.897676}, {-77.009051, 38.889939}}
		require.Condition(func() bool { return cmp.Equal(expected, line.Vertices()) })

		bin, err := json.Marshal(g)
		require.NoError(err)

		require.True(bytes.HasPrefix(bin, []byte(`{`)))
		require.True(bytes.HasSuffix(bin, []byte(`}`)))

		var g2 estype.Geoshape

		err = json.Unmarshal(bin, &g2)
		require.NoError(err)

		line2 := g2.Geometry.(geom.LineString)
		require.Condition(
			func() bool { return cmp.Equal(line.Vertices(), line2.Vertices()) },
		)
	}

	polygons := []struct {
		bin      []byte
		expected [][][2]float64
	}{
		{
			bin: []byte(`{
			"type" : "Polygon",
			"coordinates" : [[
					 [100.0, 0.0], [101.0, 0.0], [101.0, 1.0], [100.0, 1.0], [100.0, 0.0]
			]]
		}`),
			expected: [][][2]float64{{{100.0, 0.0}, {101.0, 0.0}, {101.0, 1.0}, {100.0, 1.0}, {100.0, 0.0}}},
		},
		{
			bin:      []byte(`"POLYGON ((100.0 0.0, 101.0 0.0, 101.0 1.0, 100.0 1.0, 100.0 0.0))"`),
			expected: [][][2]float64{{{100.0, 0.0}, {101.0, 0.0}, {101.0, 1.0}, {100.0, 1.0}}},
		},
		{
			bin: []byte(`{
    		"type" : "Polygon",
    		"coordinates" : [
     			[ [100.0, 0.0], [101.0, 0.0], [101.0, 1.0], [100.0, 1.0], [100.0, 0.0] ],
      			[ [100.2, 0.2], [100.8, 0.2], [100.8, 0.8], [100.2, 0.8], [100.2, 0.2] ]
    		]
  		}`),
			expected: [][][2]float64{
				{{100.0, 0.0}, {101.0, 0.0}, {101.0, 1.0}, {100.0, 1.0}, {100.0, 0.0}},
				{{100.2, 0.2}, {100.8, 0.2}, {100.8, 0.8}, {100.2, 0.8}, {100.2, 0.2}},
			},
		},

		{
			bin: []byte(`"POLYGON ((100.0 0.0, 101.0 0.0, 101.0 1.0, 100.0 1.0, 100.0 0.0),` +
				`(100.2 0.2, 100.8 0.2, 100.8 0.8, 100.2 0.8, 100.2 0.2))"`),
			expected: [][][2]float64{
				{{100.0, 0.0}, {101.0, 0.0}, {101.0, 1.0}, {100.0, 1.0}},
				{{100.2, 0.2}, {100.8, 0.2}, {100.8, 0.8}, {100.2, 0.8}},
			},
		},
	}

	for _, testCase := range polygons {
		var g estype.Geoshape

		err := json.Unmarshal(testCase.bin, &g)
		require.NoError(err)

		polygon := g.Geometry.(geom.Polygon)
		rings := polygon.LinearRings()

		require.Conditionf(
			func() bool { return cmp.Equal(testCase.expected, rings) },
			"diff: %s.\nexpected: %s\nactual: %s",
			cmp.Diff(testCase.expected, rings),
			testCase.expected,
			rings,
		)

		bin, err := json.Marshal(g)
		require.NoError(err)

		require.True(bytes.HasPrefix(bin, []byte(`{`)))
		require.True(bytes.HasSuffix(bin, []byte(`}`)))

		var g2 estype.Geoshape

		err = json.Unmarshal(bin, &g2)
		require.NoError(err)

		polygon2 := g2.Geometry.(geom.Polygon)
		require.Condition(
			func() bool { return cmp.Equal(polygon.LinearRings(), polygon2.LinearRings()) },
		)
	}
}
