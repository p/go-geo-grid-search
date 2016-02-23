package ggsearch

type locatableMap struct {
	lat_tiles int
	lng_tiles int

	locatable_map map[int]*gridTile
}

func newLocatableMap(lat_tiles, lng_tiles int) locatableMap {
	var locatable_map locatableMap
	locatable_map.lat_tiles = lat_tiles
	locatable_map.lng_tiles = lng_tiles
	locatable_map.locatable_map = make(map[int]*gridTile)
	return locatable_map
}

func (s locatableMap) AddLocatableOnGrid(locatable_on_grid *locatableOnGrid) {
	grid_pos := newGridPos(locatable_on_grid.grid_lat, locatable_on_grid.grid_lng, s.lat_tiles, s.lng_tiles)
	grid_tile := s.locatable_map[grid_pos.index]
	if grid_tile == nil {
		new_grid_tile := newGridTile(locatable_on_grid.rad_lat, locatable_on_grid.rad_lng,
			s.lat_tiles, s.lng_tiles)
		grid_tile = &new_grid_tile
		s.locatable_map[grid_pos.index] = grid_tile
	}
	grid_tile.AddLocatableOnGrid(locatable_on_grid)
}
