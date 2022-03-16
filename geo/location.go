package geo

import (
	"math"
)

const (
	mapDegreeUnit float64 = 128 / math.Pi
	pi4           float64 = math.Pi / 4
	lonEdge       float64 = 180
	latEdge       float64 = 85.05112877980659
)

type LatLon struct {
	Lat, Lon float64
}

func (loc LatLon) GetPointAtZoom(zoom int) Point {
	// Calculated based on https://en.wikipedia.org/wiki/Web_Mercator_projection
	zoomScale := 1 << zoom
	scale := mapDegreeUnit * float64(zoomScale)
	x := scale * (radians(loc.bLon()) + math.Pi)
	y := scale * (math.Pi - math.Log(math.Tan(pi4+(radians(loc.bLat())/2))))

	return Point{x, y}
}

// - Private functions

func (loc LatLon) bLon() float64 { return bounds(loc.Lon, -lonEdge, lonEdge) }

func (loc LatLon) bLat() float64 { return bounds(loc.Lat, -latEdge, latEdge) }

func radians(angle float64) float64 { return angle * math.Pi / 180 }

// Keep the `angle` within `min` and `max` edges
func bounds(angle float64, min float64, max float64) float64 {
	if angle <= min {
		return min
	} else if angle >= max {
		return max
	}

	return angle
}
