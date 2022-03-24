package geo

import (
	"testing"
)

func TestNewPoint(t *testing.T) {
	point := Point{256, 256, 0}
	if point.X != 256 && point.Y != 256 && point.Z != 0 {
		t.Error("Wrong X, Y, Z attribute values after creating a Point object")
	}
}

func TestConvertPointToLonLat(t *testing.T) {
	point := Point{X: 256, Y: 0, Z: 0}
	lonLat := point.ToLonLat()

	if lonLat.Lon != WorldBBox.Right && lonLat.Lat != WorldBBox.Top {
		t.Error("Wrong Longitude, Latitude attribute values after converting Point to LonLat")
	}
}

func TestConvertPointToTile(t *testing.T) {
	point := Point{X: 256 + 128, Y: 256 + 128, Z: 1}
	tile := point.ToTile()

	if tile.Row != 1 && tile.Column != 1 {
		t.Error("Wrong Row, Column attribute values after converting Point to Tile, values should be 1, 1")
	}
}
