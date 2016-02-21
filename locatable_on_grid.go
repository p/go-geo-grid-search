package ggsearch

type LocatableOnGrid struct {
	locatable Locatable

	rad_lat           float64
	rad_lng           float64
	grid_lat          int
	grid_lng          int
}

func NewLocatableOnGrid(locatable Locatable, lat_tiles int, lng_tiles int) LocatableOnGrid {
	locatable_on_grid := LocatableOnGrid{}
	locatable_on_grid.locatable = locatable
	locatable_on_grid.rad_lat = DegreesToRadians(locatable.Lat())
	locatable_on_grid.rad_lng = DegreesToRadians(locatable.Lng())
	locatable_on_grid.grid_lat = RadLatToGrid(locatable_on_grid.rad_lat, lat_tiles)
	locatable_on_grid.grid_lng = RadLngToGrid(locatable_on_grid.rad_lng, lng_tiles)
	return locatable_on_grid
}
