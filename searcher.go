package ggsearch

import (
	"math"
	"sort"
)

/*
Type of filtering callbacks.
*/
type Filter func(Locatable) bool

/*
Searcher maintains the grid index over Locatables in a data set and
performs K-closest queries.
*/
type Searcher struct {
	locatable_map locatableMap

	lat_tiles int
	lng_tiles int
}

/*
Creates a new Searcher with the specified number of horizontal and
vertical tiles.
*/
func NewSearcher(lat_tiles, lng_tiles int) Searcher {
	searcher := Searcher{}
	searcher.lat_tiles = lat_tiles
	searcher.lng_tiles = lng_tiles
	searcher.locatable_map = newLocatableMap(lat_tiles, lng_tiles)
	return searcher
}

/*
Adds an object satisfying the Locatable interface to this Searcher.
*/
func (s Searcher) AddLocatable(locatable Locatable) {
	locatable_on_grid := newLocatableOnGrid(locatable, s.lat_tiles, s.lng_tiles)
	s.locatable_map.AddLocatableOnGrid(&locatable_on_grid)
}

/*
Performs a K-closest search around the point identified by lat and lng
which must be given in degrees.
Limit specifies how many objects to return.
Filter is optional and if given, will be invoked for each object to
determine if the object should be in the result set.
*/
func (s Searcher) Search(filter *Filter, lat, lng float64, limit int) []Result {
	rad_lat := degreesToRadians(lat)
	rad_lng := degreesToRadians(lng)
	start_grid_lat := radLatToGrid(rad_lat, s.lat_tiles)
	start_grid_lng := radLngToGrid(rad_lng, s.lng_tiles)
	start_tile := newGridTile(rad_lat, rad_lng, s.lat_tiles, s.lng_tiles)
	tile_delta_miles := math.Min(
		start_tile.width_miles, start_tile.height_miles)

	current_list := make([]Result, 0)
	next_list := make([]Result, 0)

	reach := 0
	for reach < 5 {
		delta_miles := float64(reach) * tile_delta_miles

		new_next_list := make([]Result, 0)
		for _, locatable_in_search := range next_list {
			if locatable_in_search.distance_miles <= delta_miles {
				current_list = append(current_list, locatable_in_search)
			} else {
				new_next_list = append(current_list, locatable_in_search)
			}
		}
		next_list = new_next_list

		cells := s.cells_at_reach(start_grid_lat, start_grid_lng, reach)
		for _, grid_pos := range cells {
			grid_lat := grid_pos.grid_lat
			grid_lng := grid_pos.grid_lng
			grid_lat = clampGridLat(grid_lat, s.lat_tiles)
			grid_lng = wrapGridLng(grid_lng, s.lng_tiles)

			// xxx encapsulation violation
			tile := s.locatable_map.locatable_map[grid_pos.index]
			if tile == nil {
				continue
			}

			for _, locatable_on_grid := range tile.locatables_on_grid {
				if filter != nil && !(*filter)(locatable_on_grid.locatable) {
					continue
				}

				local_locatable_on_grid := locatable_on_grid
				locatable_in_search := newLocatableInSearch(&local_locatable_on_grid, rad_lat, rad_lng)
				if locatable_in_search.distance_miles <= delta_miles {
					current_list = append(current_list, locatable_in_search)
				} else {
					next_list = append(next_list, locatable_in_search)
				}
			}
		}

		if len(current_list) >= limit {
			break
		}

		reach++
	}

	sorted_list := locatableInSearchByDistance{current_list}
	sort.Sort(sorted_list)
	if limit > len(sorted_list.locatables_in_search) {
		limit = len(sorted_list.locatables_in_search)
	}
	locatables_in_search := sorted_list.locatables_in_search[0:limit]

	return locatables_in_search
}

func (s Searcher) cells_at_reach(grid_lat, grid_lng, reach int) []gridPos {
	if reach == 0 {
		cells := make([]gridPos, 1)
		cells[0] = newGridPos(grid_lat, grid_lng, s.lat_tiles, s.lng_tiles)
		return cells
	}

	cells := make([]gridPos, 0)
	current_grid_lat := grid_lat - reach
	current_grid_lng := grid_lng - reach
	for current_grid_lat < grid_lat+reach {
		cells = append(cells, newGridPos(current_grid_lat, current_grid_lng, s.lat_tiles, s.lng_tiles))
		current_grid_lat += 1
	}
	for current_grid_lng < grid_lng+reach {
		cells = append(cells, newGridPos(current_grid_lat, current_grid_lng, s.lat_tiles, s.lng_tiles))
		current_grid_lng += 1
	}
	for current_grid_lat > grid_lat-reach {
		cells = append(cells, newGridPos(current_grid_lat, current_grid_lng, s.lat_tiles, s.lng_tiles))
		current_grid_lat -= 1
	}
	for current_grid_lng > grid_lng-reach {
		cells = append(cells, newGridPos(current_grid_lat, current_grid_lng, s.lat_tiles, s.lng_tiles))
		current_grid_lng -= 1
	}
	return cells
}
