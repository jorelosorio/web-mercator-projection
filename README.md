# Web Mercator Projection

[![Tests](https://github.com/jorelosorio/web-mercator-projection/actions/workflows/tests.yml/badge.svg)](https://github.com/jorelosorio/web-mercator-projection/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/jorelosorio/web-mercator-projection)](https://goreportcard.com/report/github.com/jorelosorio/web-mercator-projection)
[![Coverage Status](https://coveralls.io/repos/github/jorelosorio/web-mercator-projection/badge.svg?branch=main)](https://coveralls.io/github/jorelosorio/web-mercator-projection?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/jorelosorio/web-mercator-projection.svg)](https://pkg.go.dev/github.com/jorelosorio/web-mercator-projection)

![WorldMap](https://github.com/jorelosorio/web-mercator-projection/blob/main/assets/cover.png?raw=true)

A Go project to explore the `Math` to calculate and present data in a map using the `Web Mercator Projection`.

> To get more information about this topic, please look at the following references: 

- https://en.wikipedia.org/wiki/Web_Mercator_projection
- https://www.maptiler.com/google-maps-coordinates-tile-bounds-projection/#3/15.00/50.00
- https://wiki.openstreetmap.org/wiki/Slippy_map_tilenames

## How does this project work?

In the main example project, there is a file `main.go` that contains the following actions available:

- Convert a `LonLat` into a `Point` coordinates and vice versa.

- A function to add a `Red-Marker` into a map image. (The coordinates must be setted manually in the code)
    > The resultant image will be stored in a folder named `data`.

This is how the `main.go` will print out the information:

    LonLat: {-74.92010258781309 11.045882360336755}
    Point: {298.8939304168872 480.38414652354516 2}
    Tile: {1 1 2}
    LonLat from Point: {-74.9201025878131 11.045882360336744}

### Structs definitions

- `LonLat`: A location in the map using angles in degrees.
- `Point`: A pixel in the map for a specific zoom or scale.
- `Tile`: Represent at which (tile/block) a Point or LonLat is.

> If you are curious about how the maps where built using the data in `OpenStreetMap` please take a look at this `Gist` https://gist.github.com/jorelosorio/7042bd27e4b2bb03865215d6a5607266

## Tools

- GoLang `1.17`
- Docker
- ImageMagick `6.9.11-60`
- Visual Studio Code `Optional!`
    > It requires a `Remote - Containers` extension. for more information please refers to: https://code.visualstudio.com/docs/remote/containers#_getting-started

## Development

This project contains a `Dockerfile` file with all required dependencies to run it using `Visual Studio Code` + `Remote - Containers` extension.
However, if you want to make it run locally in your development machine, please follow the instructions below.

### Install Go

Install it from https://go.dev/dl/

### Install ImageMagick

Install it using the following instructions https://imagemagick.org/script/download.php

### Build the executable

    go build ./examples/main.go

> The binary file depends on the `assets` folder to build the map, make sure to have it on the same directory.

## Licences and assets

- `OpenStreetMap` data is licenced under the Open Data Commons Open Database Licence (ODbL). https://opendatacommons.org/licenses/odbl/1-0/

- `Red marker icon` https://www.iconfinder.com/icons/299087/marker_map_icon Creative Commons (Attribution 3.0 Unported) http://creativecommons.org/licenses/by/3.0/