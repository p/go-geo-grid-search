package ggsearch

import (
	"github.com/stretchr/testify/assert" 
	"testing"
	"math"
)

func TestNewTilePos(t *testing.T) {
	grid_tile := NewGridTile(50*math.Pi/180, -70*math.Pi/180, 10, 20)
	assert.Equal(t, grid_tile.rad_lat, 50*math.Pi/180)
	assert.Equal(t, grid_tile.rad_lng, -70*math.Pi/180)
	assert.Equal(t, grid_tile.lat_tiles, 10)
	assert.Equal(t, grid_tile.lng_tiles, 20)
	// we use "spherical law of cosines" formula which produces
	// slightly different results from full haversine
	assert.InDelta(t, grid_tile.width_miles, 797, 3)
	assert.InDelta(t, grid_tile.height_miles, 1243.7, 0.1)
}

func TestAddLocatableOnGrid(t *testing.T) {
	grid_tile := NewGridTile(50*math.Pi/180, -70*math.Pi/180, 10, 20)
	assert.Equal(t, grid_tile.locatables_on_grid, make([]LocatableOnGrid, 0))
	
	locatable1 := SampleLocatable{"test1", 50, -70}
	locatable_on_grid1 := NewLocatableOnGrid(
		locatable1, 10, 10)
	grid_tile.AddLocatableOnGrid(&locatable_on_grid1)
	assert.Equal(t, len(grid_tile.locatables_on_grid), 1)
	assert.Equal(t, grid_tile.locatables_on_grid[0], locatable_on_grid1)
	
	locatable2 := SampleLocatable{"test2", 50, -71}
	locatable_on_grid2 := NewLocatableOnGrid(
		locatable2, 10, 10)
	grid_tile.AddLocatableOnGrid(&locatable_on_grid2)
	assert.Equal(t, len(grid_tile.locatables_on_grid), 2)
	assert.Equal(t, grid_tile.locatables_on_grid[0], locatable_on_grid1)
	assert.Equal(t, grid_tile.locatables_on_grid[1], locatable_on_grid2)
}
