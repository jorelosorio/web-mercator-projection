package geo

// Tile object represents tiles that conform to a map.
// A map is composite by sets of squared images of size 256x256 for the web Mercator projection,
// the tile sizes increase by the zoom level in multiples of two.
type Tile struct {
	Row, Column int
	Z           int
}
