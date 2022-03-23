package main

import (
	"fmt"

	"github.com/jorelosorio/web-mercator-projection/geo"
	"github.com/jorelosorio/web-mercator-projection/internal"
)

const zoomLevel = 2

func main() {
	location := geo.LonLat{Lon: -74.92010258781309, Lat: 11.045882360336755}
	point := location.ToPointAtZoom(zoomLevel)
	calculatedLatLon := point.ToLonLat()
	tile := point.ToTile()

	fmt.Println("LonLat:", location)
	fmt.Println("Point:", point)
	fmt.Println("Tile:", tile)
	fmt.Println("LonLat from Point:", calculatedLatLon)

	internal.CreateMarkerInMapAt(location, zoomLevel)
}
