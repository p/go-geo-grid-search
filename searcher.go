package ggsearch

import (
	"math"
	"sort"
)

type Searcher struct {
	locatable_map LocatableMap

	lat_tiles int
	lng_tiles int
}

func NewSearcher(lat_tiles, lng_tiles int) Searcher {
	searcher := Searcher{}
	searcher.lat_tiles = lat_tiles
	searcher.lng_tiles = lng_tiles
	searcher.locatable_map = NewLocatableMap(lat_tiles, lng_tiles)
	return searcher
}

func (s Searcher) AddLocatable(locatable Locatable) {
	locatable_on_grid := NewLocatableOnGrid(locatable, s.lat_tiles, s.lng_tiles)
	s.locatable_map.AddLocatableOnGrid(&locatable_on_grid)
}

type LocatableFilter func(Locatable) bool

func (s Searcher) Search(filter *LocatableFilter, lat, lng float64, limit int) []LocatableInSearch {
	rad_lat := DegreesToRadians(lat)
	rad_lng := DegreesToRadians(lng)
	start_grid_lat := RadLatToGrid(rad_lat, s.lat_tiles)
	start_grid_lng := RadLngToGrid(rad_lng, s.lng_tiles)
	start_tile := NewGridTile(rad_lat, rad_lng, s.lat_tiles, s.lng_tiles)
	tile_delta_miles := math.Min(
		start_tile.width_miles, start_tile.height_miles)

	current_list := make([]LocatableInSearch, 0)
	next_list := make([]LocatableInSearch, 0)

	reach := 0
	for reach < 5 {
		delta_miles := float64(reach) * tile_delta_miles

		new_next_list := make([]LocatableInSearch, 0)
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
			grid_lat = ClampGridLat(grid_lat, s.lat_tiles)
			grid_lng = WrapGridLng(grid_lng, s.lng_tiles)

			// xxx encapsulation violation
			tile := s.locatable_map.locatable_map[grid_pos.index]
			if tile == nil {
				continue
			}

			for _, locatable_on_grid := range tile.locatables_on_grid {
				if filter != nil && !(*filter)(locatable_on_grid.locatable) {
					continue
				}

				locatable_in_search := NewLocatableInSearch(&locatable_on_grid, rad_lat, rad_lng)
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

	sorted_list := LocatableInSearchByDistance{current_list}
	sort.Sort(sorted_list)
	if limit > len(sorted_list.locatables_in_search) {
		limit = len(sorted_list.locatables_in_search)
	}
	locatables_in_search := sorted_list.locatables_in_search[0:limit]

	return locatables_in_search
}

func (s Searcher) cells_at_reach(grid_lat, grid_lng, reach int) []GridPos {
	if reach == 0 {
		cells := make([]GridPos, 1)
		cells[0] = NewGridPos(grid_lat, grid_lng, s.lat_tiles, s.lng_tiles)
		return cells
	}

	cells := make([]GridPos, 0)
	current_grid_lat := grid_lat - reach
	current_grid_lng := grid_lng - reach
	for current_grid_lat < grid_lat+reach {
		cells = append(cells, NewGridPos(current_grid_lat, current_grid_lng, s.lat_tiles, s.lng_tiles))
		current_grid_lat += 1
	}
	for current_grid_lng < grid_lng+reach {
		cells = append(cells, NewGridPos(current_grid_lat, current_grid_lng, s.lat_tiles, s.lng_tiles))
		current_grid_lng += 1
	}
	for current_grid_lat > grid_lat-reach {
		cells = append(cells, NewGridPos(current_grid_lat, current_grid_lng, s.lat_tiles, s.lng_tiles))
		current_grid_lat -= 1
	}
	for current_grid_lng > grid_lng-reach {
		cells = append(cells, NewGridPos(current_grid_lat, current_grid_lng, s.lat_tiles, s.lng_tiles))
		current_grid_lng -= 1
	}
	return cells
}
