package geo

import (
	"fmt"
	"testing"
)

func TestNewLonLat(t *testing.T) {
	lonLat := LonLat{Lon: -74.92010259, Lat: 11.04588236}
	if lonLat.Lon != -74.92010259 && lonLat.Lat != 11.04588236 {
		t.Error("Wrong Longitude, Latitude attribute values after creating a LonLat object")
	}
}

func TestConvertToString(t *testing.T) {
	lonLat := LonLat{Lon: WorldBBox.Right, Lat: WorldBBox.Top}

	if fmt.Sprint(lonLat) != fmt.Sprintf("Longitude: %.8f, Latitude: %.8f", lonLat.Lon, lonLat.Lat) {
		t.Error("Invalid String format when printing LonLat attributes")
	}
}

func TestConvertToPointExceedsTopLeftEdges(t *testing.T) {
	lonLat := LonLat{Lon: -200, Lat: -90}
	point := lonLat.ToPointAtZoom(0)

	if point.X != 0 && point.Y != 0 && point.Z != 1 {
		t.Error("Wrong X, Y, Z attribute values after converting LonLat to Point")
	}
}

func TestConvertToPointExceedsRightBottomEdges(t *testing.T) {
	lonLat := LonLat{Lon: 200, Lat: 90}
	point := lonLat.ToPointAtZoom(0)

	if point.X != 256 && point.Y != 256 && point.Z != 1 {
		t.Error("Wrong X, Y, Z attribute values after converting LonLat to Point")
	}
}

func TestConvertToPointTopLeft(t *testing.T) {
	lonLat := LonLat{Lon: WorldBBox.Left, Lat: WorldBBox.Bottom}
	point := lonLat.ToPointAtZoom(1)

	if point.X != 0 && point.Y != 0 && point.Z != 1 {
		t.Error("Wrong X, Y, Z attribute values after converting LonLat to Point")
	}
}

func TestConvertToPointRightBottom(t *testing.T) {
	lonLat := LonLat{Lon: WorldBBox.Right, Lat: WorldBBox.Top}
	point := lonLat.ToPointAtZoom(1)

	if point.X != 512 && point.Y != 512 && point.Z != 1 {
		t.Error("Wrong X, Y, Z attribute values after converting LonLat to Point")
	}
}
