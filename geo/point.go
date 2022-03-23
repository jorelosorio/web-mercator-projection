package geo

import "math"

// Point object represents a position in a map image in pixels, for a given zoom level.
type Point struct {
	X, Y float64
	Z    int
}

// ToLonLat converts a pixel Point to a LonLat object.
func (p Point) ToLonLat() LonLat {
	gridSize := float64(p.numberOfTiles()) * TileSizeInRadians

	n := math.Pi - (p.Y / gridSize)
	lon := (p.X / gridSize) - math.Pi
	lat := math.Atan(0.5 * (math.Exp(n) - math.Exp(-n)))

	return LonLat{Lon: RadToDeg(lon), Lat: RadToDeg(lat)}
}

// ToTile converts a pixel point to a Tile object.
func (p Point) ToTile() Tile {
	row := int(math.Max(0, math.Ceil(p.X/TileSizeInPixels)-1))
	column := int(math.Min(float64(p.numberOfTiles()-1), math.Ceil(p.Y/TileSizeInPixels)) - 1)

	return Tile{row, column, p.Z}
}

func (p Point) numberOfTiles() int {
	return 1 << p.Z
}
