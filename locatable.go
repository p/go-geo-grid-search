package ggsearch

type Locatable interface {
	Lat() float64
	Lng() float64
}
