package ggsearch

import (
	"github.com/stretchr/testify/assert"
	check "gopkg.in/check.v1"
	"math"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func TestGeoUtils(t *testing.T) { check.TestingT(t) }

type GeoUtilsSuite struct {
}

var _ = check.Suite(&GeoUtilsSuite{})

func (s *GeoUtilsSuite) TestDegreesToRadiansZero(c *check.C) {
	c.Assert(degreesToRadians(0), check.Equals, 0.0)
}

func (s *GeoUtilsSuite) TestDegreesToRadiansHalfPi(c *check.C) {
	c.Assert(degreesToRadians(90), check.Equals, math.Pi/2)
}

func (s *GeoUtilsSuite) TestDegreesToRadiansPi(c *check.C) {
	c.Assert(degreesToRadians(180), check.Equals, math.Pi)
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

func (s *GeoUtilsSuite) TestClampGridLat(c *check.C) {
	c.Assert(clampGridLat(10, 11), check.Equals, 10)
	c.Assert(clampGridLat(11, 11), check.Equals, 10)
	c.Assert(clampGridLat(0, 11), check.Equals, 0)
	c.Assert(clampGridLat(-1, 11), check.Equals, 0)
}

func (s *GeoUtilsSuite) TestWrapGridLng(c *check.C) {
	c.Assert(wrapGridLng(10, 11), check.Equals, 10)
	c.Assert(wrapGridLng(11, 11), check.Equals, 0)
	c.Assert(wrapGridLng(0, 11), check.Equals, 0)
	c.Assert(wrapGridLng(-1, 11), check.Equals, 10)
}

func TestHaversine(t *testing.T) {
	assert := assert.New(t)
	distance := haversine(50*math.Pi/180, -70*math.Pi/180,
		51*math.Pi/180, -71*math.Pi/180)
	assert.InDelta(distance, 81.89, 0.01)
}
