package ggsearch

import (
	"github.com/stretchr/testify/assert"
	. "gopkg.in/check.v1"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func TestLocatableOnGrid(t *testing.T) { TestingT(t) }

type LocatableOnGridSuite struct {
}

var _ = Suite(&LocatableOnGridSuite{})

func TestGridIndexCalculation(t *testing.T) {
	locatable := SampleLocatable{"", 50, -70}
	locatable_on_grid := newLocatableOnGrid(
		locatable, 10, 10)
	assert.Equal(t, locatable_on_grid.locatable, locatable)
	assert.InDelta(t, locatable_on_grid.rad_lat, 0.87, 0.01)
	assert.InDelta(t, locatable_on_grid.rad_lng, -1.22, 0.01)
	assert.Equal(t, locatable_on_grid.grid_lat, 7)
	assert.Equal(t, locatable_on_grid.grid_lng, 3)
}
