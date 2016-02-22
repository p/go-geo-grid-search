package ggsearch

type LocatableMap struct {
	lat_tiles int
	lng_tiles int

	locatable_map map[int]*GridTile
}

func NewLocatableMap(lat_tiles, lng_tiles int) LocatableMap {
	var locatable_map LocatableMap
	locatable_map.lat_tiles = lat_tiles
	locatable_map.lng_tiles = lng_tiles
	locatable_map.locatable_map = make(map[int]*GridTile)
	return locatable_map
}

func (s LocatableMap) AddLocatableOnGrid(locatable_on_grid *LocatableOnGrid) {
	grid_pos := NewGridPos(locatable_on_grid.grid_lat, locatable_on_grid.grid_lng, s.lat_tiles, s.lng_tiles)
	grid_tile := s.locatable_map[grid_pos.index]
	if grid_tile == nil {
		new_grid_tile := NewGridTile(locatable_on_grid.rad_lat, locatable_on_grid.rad_lng,
			s.lat_tiles, s.lng_tiles)
		grid_tile = &new_grid_tile
		s.locatable_map[grid_pos.index] = grid_tile
	}
	grid_tile.AddLocatableOnGrid(locatable_on_grid)
}
