package geo

import (
	"fmt"
	"math"
)

// LonLat object represents a position in degrees for the Mercator projection.
type LonLat struct {
	Lon, Lat float64
}

// ToPointAtZoom converts a LonLat object to pixel Point at a specified zoom level.
func (l LonLat) ToPointAtZoom(zoomLevel int) Point {
	// Calculated based on https://en.wikipedia.org/wiki/Web_Mercator_projection
	numberOfTiles := 1 << zoomLevel
	mapSize := float64(numberOfTiles) * TileSizeInRadians
	x := mapSize * (DegToRad(l.fixedLon()) + math.Pi)
	y := mapSize * (math.Pi - math.Log(math.Tan(Pi4+(DegToRad(l.fixedLat())/2))))

	return Point{X: x, Y: y, Z: zoomLevel}
}

// String prints the longitude, latitude degrees
func (l LonLat) String() string {
	return fmt.Sprintf("Longitude: %.8f, Latitude: %.8f", l.Lon, l.Lat)
}

func (l LonLat) fixedLon() float64 { return WorldBBox.WithinXAxis(l.Lon) }

func (l LonLat) fixedLat() float64 { return WorldBBox.WithinYAxis(l.Lat) }
