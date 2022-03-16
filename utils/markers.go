package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"web-mercator/geo"
)

func CreateMarkerInMapAt(location geo.LonLat, zoomLevel int) {
	point := location.ToPointAtZoom(zoomLevel)

	// Marker position must be aligned to the centre, bottom
	markerPosX, markerPosY := point.X-16, point.Y-48

	// Assets paths
	currentPath, _ := os.Getwd()
	imagesPath := filepath.Join(currentPath, "assets")
	markerIconPath := filepath.Join(imagesPath, "red_marker_map_icon.png")
	mapPath := filepath.Join(imagesPath, "maps", fmt.Sprintf("open_street_map_zoom_lvl_%d.png", zoomLevel))

	drawMarkerArgs := fmt.Sprintf("image SrcOver %f,%f 0,0 \"%s\"", markerPosX, markerPosY, markerIconPath)
	exportPath := filepath.Join(currentPath, "data")

	_ = os.Mkdir(exportPath, os.ModeDir)

	// Run ImageMagick command:
	// It creates a new map image with a marker icon on it
	cmd := exec.Command("mogrify", "-path", exportPath, "-draw", drawMarkerArgs, mapPath)
	cmd.Run()
}
