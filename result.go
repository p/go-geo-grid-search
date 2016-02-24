package ggsearch

/*
Search result, provides access to the original Locatable object
and distance from query point to this object.
*/
type Result struct {
	locatable_on_grid *locatableOnGrid

	query_rad_lat  float64
	query_rad_lng  float64
	distance_miles float64
}

func newLocatableInSearch(locatable_on_grid *locatableOnGrid, query_rad_lat, query_rad_lng float64) Result {
	locatable_in_search := Result{}
	locatable_in_search.locatable_on_grid = locatable_on_grid
	locatable_in_search.query_rad_lat = query_rad_lat
	locatable_in_search.query_rad_lng = query_rad_lng

	locatable_in_search.distance_miles = haversine(
		locatable_on_grid.rad_lat, locatable_on_grid.rad_lng, query_rad_lat, query_rad_lng)
	return locatable_in_search
}

/*
Returns the Locatable object for this search result.
*/
func (s Result) Locatable() Locatable {
	return s.locatable_on_grid.locatable
}

/*
Returns the distance, in miles, from the Locatable object in this search
result to the query point.
*/
func (s Result) Distance() float64 {
	return s.distance_miles
}

type locatableInSearchByDistance struct {
	locatables_in_search []Result
}

func (s locatableInSearchByDistance) Len() int {
	return len(s.locatables_in_search)
}

func (s locatableInSearchByDistance) Swap(i, j int) {
	s.locatables_in_search[i], s.locatables_in_search[j] =
		s.locatables_in_search[j], s.locatables_in_search[i]
}

func (s locatableInSearchByDistance) Less(i, j int) bool {
	return s.locatables_in_search[i].distance_miles < s.locatables_in_search[j].distance_miles
}
