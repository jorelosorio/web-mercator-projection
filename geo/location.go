package geo

import (
	"math"
)

type LonLat struct {
	Lon, Lat float64
}

func (loc LonLat) ToPointAtZoom(zoomLevel int) Point {
	// Calculated based on https://en.wikipedia.org/wiki/Web_Mercator_projection
	gridSize := 1 << zoomLevel
	mapSize := TileSizeInRadians * float64(gridSize)
	x := mapSize * (ToRadians(loc.fixedLon()) + math.Pi)
	y := mapSize * (math.Pi - math.Log(math.Tan(pi4+(ToRadians(loc.fixedLat())/2))))

	return Point{X: x, Y: y, Z: zoomLevel}
}

func (lonLat LonLat) fixedLon() float64 { return withinMinMax(lonLat.Lon, -lonEdge, lonEdge) }

func (lonLat LonLat) fixedLat() float64 { return withinMinMax(lonLat.Lat, -latEdge, latEdge) }
