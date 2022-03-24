package geo

import "math"

const (
	//
	// Convertions
	//
	radians float64 = math.Pi / 180.0
	degrees float64 = 180.0 / math.Pi
	twoPi   float64 = math.Pi * 2
	// Pi4 is a quarter of Pi
	Pi4 float64 = math.Pi / 4
	//
	// Web Mercator projection constraints
	//
	// lonEdge edges from −180° to 180° Longitude
	lonEdge float64 = 180
	// latEdge edges from -85.05° North and 85.05° South (Also know as yMax)
	latEdge float64 = 85.05112877980659
)

const (
	// TileSizeInPixels size is 256x256
	TileSizeInPixels float64 = 256.0
	// TileSizeInRadians a value that represents a map that extends from 0 to 256 / (2 * PI)
	TileSizeInRadians float64 = 128 / math.Pi
)

var (
	// WorldBBox contains the web mercator edges values
	WorldBBox BBox = BBox{Left: -lonEdge, Right: lonEdge, Top: -latEdge, Bottom: latEdge}
)
