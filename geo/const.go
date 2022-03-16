package geo

import "math"

const (
	//
	// Convertions
	//
	radians float64 = math.Pi / 180.0
	degrees float64 = 180.0 / math.Pi
	twoPi   float64 = math.Pi * 2
	pi4     float64 = math.Pi / 4
	//
	// Map related
	//
	TileSizeInPixels  int     = 256
	TileSizeInRadians float64 = 128 / math.Pi // A value that represents a map that extends from 0 to 256 / (2 * PI)
	lonEdge           float64 = 180
	latEdge           float64 = 85.05112877980659 // Also know as yMax
)
