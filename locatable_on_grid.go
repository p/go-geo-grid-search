package ggsearch

type locatableOnGrid struct {
	locatable Locatable

	rad_lat  float64
	rad_lng  float64
	grid_lat int
	grid_lng int
}

func newLocatableOnGrid(locatable Locatable, lat_tiles int, lng_tiles int) locatableOnGrid {
	locatable_on_grid := locatableOnGrid{}
	locatable_on_grid.locatable = locatable
	locatable_on_grid.rad_lat = degreesToRadians(locatable.Lat())
	locatable_on_grid.rad_lng = degreesToRadians(locatable.Lng())
	locatable_on_grid.grid_lat = radLatToGrid(locatable_on_grid.rad_lat, lat_tiles)
	locatable_on_grid.grid_lng = radLngToGrid(locatable_on_grid.rad_lng, lng_tiles)
	return locatable_on_grid
}
