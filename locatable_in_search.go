package ggsearch

type LocatableInSearch struct {
	locatable_on_grid *locatableOnGrid

	query_rad_lat  float64
	query_rad_lng  float64
	distance_miles float64
}

func newLocatableInSearch(locatable_on_grid *locatableOnGrid, query_rad_lat, query_rad_lng float64) LocatableInSearch {
	locatable_in_search := LocatableInSearch{}
	locatable_in_search.locatable_on_grid = locatable_on_grid
	locatable_in_search.query_rad_lat = query_rad_lat
	locatable_in_search.query_rad_lng = query_rad_lng

	locatable_in_search.distance_miles = haversine(
		locatable_on_grid.rad_lat, locatable_on_grid.rad_lng, query_rad_lat, query_rad_lng)
	return locatable_in_search
}

func (s LocatableInSearch) GetLocatable() Locatable {
	return s.locatable_on_grid.locatable
}

type LocatableInSearchByDistance struct {
	locatables_in_search []LocatableInSearch
}

func (s LocatableInSearchByDistance) Len() int {
	return len(s.locatables_in_search)
}

func (s LocatableInSearchByDistance) Swap(i, j int) {
	s.locatables_in_search[i], s.locatables_in_search[j] =
		s.locatables_in_search[j], s.locatables_in_search[i]
}

func (s LocatableInSearchByDistance) Less(i, j int) bool {
	return s.locatables_in_search[i].distance_miles < s.locatables_in_search[j].distance_miles
}
