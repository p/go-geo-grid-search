package ggsearch

import (
	"math"
	. "gopkg.in/check.v1"
	"testing"
	"github.com/stretchr/testify/assert" 
)

// Hook up gocheck into the "go test" runner.
func TestGeoUtils(t *testing.T) { TestingT(t) }

type GeoUtilsSuite struct {
}

var _ = Suite(&GeoUtilsSuite{})

func (s *GeoUtilsSuite) TestDegreesToRadiansZero(c *C) {
	c.Assert(DegreesToRadians(0), Equals, 0.0)
}

func (s *GeoUtilsSuite) TestDegreesToRadiansHalfPi(c *C) {
	c.Assert(DegreesToRadians(90), Equals, math.Pi/2)
}

func (s *GeoUtilsSuite) TestDegreesToRadiansPi(c *C) {
	c.Assert(DegreesToRadians(180), Equals, math.Pi)
}

func TestRadLatToGrid(t *testing.T) {
	assert.Equal(t, RadLatToGrid(0, 10), 5)
	assert.Equal(t, RadLatToGrid(-0.1, 10), 4)
	assert.Equal(t, RadLatToGrid(math.Pi/2-0.1, 10), 9)
	assert.Equal(t, RadLatToGrid(-math.Pi/2+0.1, 10), 0)
}

func TestRadLngToGrid(t *testing.T) {
	assert.Equal(t, RadLngToGrid(0, 10), 5)
	assert.Equal(t, RadLngToGrid(-0.1, 10), 4)
	assert.Equal(t, RadLngToGrid(math.Pi-0.1, 10), 9)
	assert.Equal(t, RadLngToGrid(-math.Pi+0.1, 10), 0)
}

func (s *GeoUtilsSuite) TestClampGridLat(c *C) {
	c.Assert(ClampGridLat(10, 11), Equals, 10)
	c.Assert(ClampGridLat(11, 11), Equals, 10)
	c.Assert(ClampGridLat(0, 11), Equals, 0)
	c.Assert(ClampGridLat(-1, 11), Equals, 0)
}

func (s *GeoUtilsSuite) TestWrapGridLng(c *C) {
	c.Assert(WrapGridLng(10, 11), Equals, 10)
	c.Assert(WrapGridLng(11, 11), Equals, 0)
	c.Assert(WrapGridLng(0, 11), Equals, 0)
	c.Assert(WrapGridLng(-1, 11), Equals, 10)
}

func TestHaversine(t *testing.T) {
	assert := assert.New(t)
	distance := Haversine(50*math.Pi/180, -70*math.Pi/180,
		51*math.Pi/180, -71*math.Pi/180)
	assert.InDelta(distance, 81.89, 0.01)
}
