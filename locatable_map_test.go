package ggsearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLocatableMap(t *testing.T) {
	locatable_map := newLocatableMap(10, 20)
	assert.Equal(t, locatable_map.lat_tiles, 10)
	assert.Equal(t, locatable_map.lng_tiles, 20)
	assert.Equal(t, locatable_map.locatable_map, make(map[int]*gridTile))
}

func TestLocatableMap_AddLocatableOnGrid(t *testing.T) {
	locatable_map := newLocatableMap(10, 20)
	assert.Equal(t, locatable_map.locatable_map, make(map[int]*gridTile))

	locatable1 := SampleLocatable{"test1", 0, 0}
	locatable_on_grid1 := newLocatableOnGrid(
		locatable1, 10, 20)
	assert.Equal(t, locatable_on_grid1.grid_lat, 5)
	assert.Equal(t, locatable_on_grid1.grid_lng, 10)
	assert.Equal(t, newGridPos(locatable_on_grid1.grid_lat,
		locatable_on_grid1.grid_lng, 10, 20).index, 105)
	locatable_map.AddLocatableOnGrid(&locatable_on_grid1)
	assert.Equal(t, len(locatable_map.locatable_map), 1)
	assert.Equal(t, locatable_map.locatable_map[105].locatables_on_grid[0], locatable_on_grid1)

	locatable2 := SampleLocatable{"test2", 0, -1}
	locatable_on_grid2 := newLocatableOnGrid(
		locatable2, 10, 20)
	assert.Equal(t, newGridPos(locatable_on_grid2.grid_lat,
		locatable_on_grid2.grid_lng, 10, 20).index, 95)
	locatable_map.AddLocatableOnGrid(&locatable_on_grid2)
	assert.Equal(t, len(locatable_map.locatable_map), 2)
	assert.Equal(t, locatable_map.locatable_map[105].locatables_on_grid[0], locatable_on_grid1)
	assert.Equal(t, locatable_map.locatable_map[95].locatables_on_grid[0], locatable_on_grid2)
}
