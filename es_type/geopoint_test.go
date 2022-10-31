package estype_test

import (
	"bytes"
	"encoding/json"
	"testing"

	estype "github.com/ngicks/elastic-type/es_type"
	"github.com/stretchr/testify/require"
)

func TestGeopoint(t *testing.T) {
	inputs := [][]byte{
		[]byte(`{"type": "Point", "coordinates": [-71.34, 41.12]}`),
		[]byte(`{"lat": 41.12, "lon": -71.34}`),
		[]byte(`[ -71.34, 41.12 ]`),
		[]byte(`"41.12,-71.34"`),
		[]byte(`"drm3btev3e86"`),
		[]byte(`"POINT (-71.34 41.12)"`),
	}

	for _, v := range inputs {
		var g estype.Geopoint

		err := json.Unmarshal(v, &g)
		require.NoError(t, err)

		require.InDelta(t, 41.12, g.Lat, 0.001)
		require.InDelta(t, -71.34, g.Lon, 0.001)

		bin, err := json.Marshal(g)
		require.NoError(t, err)

		require.True(t, bytes.Contains(bin, []byte(`lat`)))
		require.True(t, bytes.Contains(bin, []byte(`lon`)))

		var g2 estype.Geopoint

		err = json.Unmarshal(bin, &g2)
		require.NoError(t, err)

		require.Equal(t, g, g2)
	}
}
