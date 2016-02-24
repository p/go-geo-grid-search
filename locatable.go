package ggsearch

/*
ggsearch stores and queries objects implementing Locatable interface.
*/
type Locatable interface {
	// Returns latitude of this object in degrees
	Lat() float64
	// Returns longitude of this object in degrees
	Lng() float64
}
