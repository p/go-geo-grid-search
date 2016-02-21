package ggsearch

import (
	"math"
)

type GridTile struct {
	rad_lat         float64
	rad_lng         float64
	lat_tiles       int
	lng_tiles       int
	locatables_on_grid []LocatableOnGrid

	width_miles  float64
	height_miles float64
}

func NewGridTile(
	rad_lat, rad_lng float64, lat_tiles, lng_tiles int) GridTile {
	var grid_tile GridTile
	grid_tile.rad_lat = rad_lat
	grid_tile.rad_lng = rad_lng
	grid_tile.lat_tiles = lat_tiles
	grid_tile.lng_tiles = lng_tiles
	grid_tile.locatables_on_grid = make([]LocatableOnGrid, 0)

	rad_lng_length := 2 * math.Pi / float64(lng_tiles)

	grid_tile.width_miles = Haversine(rad_lat, rad_lng-rad_lng_length/2,
		rad_lat, rad_lng+rad_lng_length/2)
	grid_tile.height_miles = R * math.Pi / float64(lat_tiles)
	return grid_tile
}

func (s *GridTile) AddLocatableOnGrid(locatable_on_grid *LocatableOnGrid) {
	s.locatables_on_grid = append(s.locatables_on_grid, *locatable_on_grid)
}
