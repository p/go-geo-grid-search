package ggsearch

type GridPos struct {
	grid_lat int
	grid_lng int

	index int
}

func NewGridPos(grid_lat, grid_lng, lat_tiles, lng_tiles int) GridPos {
	var grid_pos GridPos
	grid_pos.grid_lat = grid_lat
	grid_pos.grid_lng = grid_lng
	grid_pos.index = grid_lat + grid_lng*lat_tiles
	return grid_pos
}
