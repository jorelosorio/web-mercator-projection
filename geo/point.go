package geo

import "math"

type Point struct {
	X, Y float64
	Z    int
}

func (p Point) ToLonLat() LonLat {
	gridSize := float64(p.numberOfTiles()) * TileSizeInRadians

	n := math.Pi - (p.Y / gridSize)
	lon := (p.X / gridSize) - math.Pi
	lat := math.Atan(0.5 * (math.Exp(n) - math.Exp(-n)))

	return LonLat{Lon: ToDegrees(lon), Lat: ToDegrees(lat)}
}

func (p Point) ToTile() Tile {
	gridSize := float64(p.numberOfTiles() * TileSizeInPixels)
	row, column := int(math.Ceil(p.X/gridSize)), int(math.Ceil(p.Y/gridSize))

	return Tile{row, column}
}

func (p Point) numberOfTiles() int {
	return 1 << p.Z
}
