package estype

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/go-spatial/geom"
	"github.com/go-spatial/geom/encoding/geojson"
	"github.com/go-spatial/geom/encoding/wkt"
)

// TODO: test against elasticsearch 8.4
type Geoshape struct {
	// avoid embedding this because it could confuse user when re-defined type is made in user code.
	Geometry geom.Geometry
}

func (g *Geoshape) UnmarshalJSON(data []byte) error {
	data = bytes.TrimLeft(data, " ")

	switch data[0] {
	case '{':
		var geo geojson.Geometry
		err := geo.UnmarshalJSON(data)
		if err != nil {
			return err
		}
		g.Geometry = geo.Geometry
		return nil
	case '"':
		geo, err := wkt.DecodeBytes(data[1 : len(data)-1])
		if err != nil {
			return err
		}
		g.Geometry = geo
		return nil
	}

	return fmt.Errorf(
		"unknown type: must be geojson or geohash string literal, but was %s",
		string(data),
	)
}

func (g Geoshape) MarshalJSON() ([]byte, error) {
	return json.Marshal(geojson.Geometry{Geometry: g.Geometry})
}
