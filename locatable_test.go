package ggsearch

type SampleLocatable struct {
	name string
	lat  float64
	lng  float64
}

func (s SampleLocatable) Lat() float64 {
	return s.lat
}

func (s SampleLocatable) Lng() float64 {
	return s.lng
}
