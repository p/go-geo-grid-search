package ggsearch

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestDegreesToRadiansZero(t *testing.T) {
	assert.Equal(t, degreesToRadians(0), 0.0)
}

func TestDegreesToRadiansHalfPi(t *testing.T) {
	assert.Equal(t, degreesToRadians(90), math.Pi/2)
}

func TestDegreesToRadiansPi(t *testing.T) {
	assert.Equal(t, degreesToRadians(180), math.Pi)
}

func TestRadLatToGrid(t *testing.T) {
	assert.Equal(t, radLatToGrid(0, 10), 5)
	assert.Equal(t, radLatToGrid(-0.1, 10), 4)
	assert.Equal(t, radLatToGrid(math.Pi/2-0.1, 10), 9)
	assert.Equal(t, radLatToGrid(-math.Pi/2+0.1, 10), 0)
}

func TestRadLngToGrid(t *testing.T) {
	assert.Equal(t, radLngToGrid(0, 10), 5)
	assert.Equal(t, radLngToGrid(-0.1, 10), 4)
	assert.Equal(t, radLngToGrid(math.Pi-0.1, 10), 9)
	assert.Equal(t, radLngToGrid(-math.Pi+0.1, 10), 0)
}

func TestClampGridLat(t *testing.T) {
	assert.Equal(t, clampGridLat(10, 11), 10)
	assert.Equal(t, clampGridLat(11, 11), 10)
	assert.Equal(t, clampGridLat(0, 11), 0)
	assert.Equal(t, clampGridLat(-1, 11), 0)
}

func TestWrapGridLng(t *testing.T) {
	assert.Equal(t, wrapGridLng(10, 11), 10)
	assert.Equal(t, wrapGridLng(11, 11), 0)
	assert.Equal(t, wrapGridLng(0, 11), 0)
	assert.Equal(t, wrapGridLng(-1, 11), 10)
}

func TestHaversine(t *testing.T) {
	assert := assert.New(t)
	distance := haversine(50*math.Pi/180, -70*math.Pi/180,
		51*math.Pi/180, -71*math.Pi/180)
	assert.InDelta(distance, 81.89, 0.01)
}
