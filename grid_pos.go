package ggsearch

type gridPos struct {
	grid_lat int
	grid_lng int

	index int
}

func newGridPos(grid_lat, grid_lng, lat_tiles, lng_tiles int) gridPos {
	var grid_pos gridPos
	grid_pos.grid_lat = grid_lat
	grid_pos.grid_lng = grid_lng
	grid_pos.index = grid_lat + grid_lng*lat_tiles
	return grid_pos
}
