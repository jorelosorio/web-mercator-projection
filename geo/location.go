package geo

import (
	"math"
)

type LonLat struct {
	Lon, Lat float64
}

func (l LonLat) ToPointAtZoom(zoomLevel int) Point {
	// Calculated based on https://en.wikipedia.org/wiki/Web_Mercator_projection
	numberOfTiles := 1 << zoomLevel
	mapSize := float64(numberOfTiles) * TileSizeInRadians
	x := mapSize * (ToRadians(l.fixedLon()) + math.Pi)
	y := mapSize * (math.Pi - math.Log(math.Tan(Pi4+(ToRadians(l.fixedLat())/2))))

	return Point{X: x, Y: y, Z: zoomLevel}
}

func (l LonLat) fixedLon() float64 { return withinMinMax(l.Lon, -LonEdge, LonEdge) }

func (l LonLat) fixedLat() float64 { return withinMinMax(l.Lat, -LatEdge, LatEdge) }
