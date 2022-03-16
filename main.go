package main

import (
	"web-mercator/geo"
	"web-mercator/utils"
)

const zoomLevel = 2

func main() {
	location := geo.LatLon{Lon: -74.92010258781309, Lat: 11.045882360336755}
	utils.CreateMarkerInMapAt(location, zoomLevel)
}
