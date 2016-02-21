package ggsearch

import (
	"github.com/stretchr/testify/assert" 
	"testing"
)

func TestNewGridPos(t *testing.T) {
	grid_pos := NewGridPos(4, 6, 10, 20)
	assert.Equal(t, grid_pos.index, 64)
}
