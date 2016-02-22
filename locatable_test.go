package ggsearch

import (
	. "gopkg.in/check.v1"
	"testing"
)

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

var _ Locatable = (*SampleLocatable)(nil)

// Hook up gocheck into the "go test" runner.
func TestLocatable(t *testing.T) { TestingT(t) }

type LocatableSuite struct {
}

var _ = Suite(&LocatableSuite{})
