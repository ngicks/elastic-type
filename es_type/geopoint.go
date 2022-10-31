package estype

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-spatial/geom"
	"github.com/go-spatial/geom/encoding/wkt"
	"github.com/mmcloughlin/geohash"
)

func UnmarshalEsGeopointJSON(data []byte) (Geopoint, error) {
	data = bytes.Trim(data, " ")
	if len(data) < 3 {
		return Geopoint{}, fmt.Errorf("too short: input must longer than 3 chars but %d", len(data))
	}
	switch data[0] {
	case '[':
		// [lon, lat]
		var d []float64
		err := json.Unmarshal(data, &d)
		if err != nil {
			return Geopoint{}, err
		}
		if len(d) < 2 {
			return Geopoint{}, fmt.Errorf("too short: must be [lon, lat] but %d", len(d))
		}

		return Geopoint{
			Lon: d[0],
			Lat: d[1],
		}, nil
	case '"':
		// geohash
		// Well-known text: POINT(lon lat)
		// lat,lon
		return UnmarshalEsGeopointText(data[1 : len(data)-1])
	case '{':
		if bytes.Contains(data, []byte("type")) {
			type simpleGeoJSON struct {
				Type        string    `json:"type"`
				Coordinates []float64 `json:"coordinates"` // see https://www.rfc-editor.org/rfc/rfc7946#section-3.1.1
			}

			var j simpleGeoJSON
			err := json.Unmarshal(data, &j)
			if err != nil {
				return Geopoint{}, err
			}

			if j.Type != "Point" {
				return Geopoint{}, fmt.Errorf(fmt.Sprintf("type must be Point but is %s", j.Type))
			}

			if len(j.Coordinates) < 2 {
				return Geopoint{}, fmt.Errorf(`too short: must have coodinate which is Number[] of at least 2 elements`)
			}

			return Geopoint{
				Lon: j.Coordinates[0],
				Lat: j.Coordinates[1],
			}, nil
		}

		if bytes.Contains(data, []byte("lat")) && bytes.Contains(data, []byte("lon")) {
			type simpleGeopoint struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			}

			var p simpleGeopoint
			err := json.Unmarshal(data, &p)
			if err != nil {
				return Geopoint{}, err
			}

			return Geopoint(p), nil
		}

		return Geopoint{}, fmt.Errorf("unknown format: " + string(data))
	}
	return Geopoint{}, fmt.Errorf("unknown format: " + string(data))
}

func UnmarshalEsGeopointText(text []byte) (Geopoint, error) {
	strText := string(text)
	if strText[0] == 'P' {
		// POINT(lon lat)
		geo, err := wkt.Decode(strings.NewReader(strText))
		if err != nil {
			return Geopoint{}, err
		}

		point, ok := geo.(geom.Point)
		if !ok {
			return Geopoint{}, fmt.Errorf(fmt.Sprintf("unknown: must be point but is %T", geo))
		}

		return Geopoint{
			Lon: point[0], // x
			Lat: point[1], // y
		}, nil
	}

	if strings.Contains(strText, ",") {
		// lat,lon
		points := strings.Split(strText, ",")
		if len(points) < 2 {
			return Geopoint{}, fmt.Errorf(`too short: must be "lat,lon", or more`)
		}

		lat, err := strconv.ParseFloat(strings.Trim(points[0], " "), 64)
		if err != nil {
			return Geopoint{}, err
		}
		lon, err := strconv.ParseFloat(strings.Trim(points[1], " "), 64)
		if err != nil {
			return Geopoint{}, err
		}

		return Geopoint{
			Lat: lat,
			Lon: lon,
		}, nil
	}

	if err := geohash.Validate(strText); err == nil {
		lat, lon := geohash.Decode(strText)

		return Geopoint{
			Lat: lat,
			Lon: lon,
		}, nil
	}

	return Geopoint{}, fmt.Errorf(
		fmt.Sprintf(
			`uknown: %s. UnmarshalText only supports `+
				`well-known text: "POINT(lon lat)", `+
				`es old text: "lat,lon" or `+
				`geohash: "drm3btev3e86"`,
			strText,
		),
	)
}

// Elasticsearch geopoint type.
//
// For historical reason, it has 6 formats to represent same data type.
// This type marshal into only one format, simple json consists of lat and lon keys populated with Number,
// namely, `{"lat":123, "lon":456}`.
//
// see: https://www.elastic.co/guide/en/elasticsearch/reference/8.4/geo-point.html
type Geopoint struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func (g *Geopoint) UnmarshalJSON(data []byte) error {
	var err error
	*g, err = UnmarshalEsGeopointJSON(data)
	if err != nil {
		return err
	}
	return nil
}

func (g *Geopoint) UnmarshalText(text []byte) error {
	var err error
	*g, err = UnmarshalEsGeopointText(text)
	if err != nil {
		return err
	}
	return nil
}
