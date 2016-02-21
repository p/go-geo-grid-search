package ggsearch

import (
	"math"
)

// Earth radius in miles
const R = 3959

func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// http://www.movable-type.co.uk/scripts/latlong.html
func Haversine(rad_lat1, rad_lng1, rad_lat2, rad_lng2 float64) float64 {
	delta_lambda := rad_lng2 - rad_lng1
	delta_phi := rad_lat2 - rad_lat1
	phi_m := (rad_lat1 + rad_lat2) / 2
	x := delta_lambda * math.Cos(phi_m)
	y := delta_phi
	hfo := x*x + y*y
	return R * math.Sqrt(hfo)
}

func RadLatToGrid(rad_lat float64, lat_tiles int) int {
	return int((rad_lat + math.Pi/2) * float64(lat_tiles) / math.Pi)
}

func RadLngToGrid(rad_lng float64, lng_tiles int) int {
	return int((rad_lng + math.Pi) * float64(lng_tiles) / math.Pi / 2)
}

func ClampGridLat(grid_lat, lat_tiles int) int {
	if grid_lat >= lat_tiles {
		grid_lat -= 1
	}
	if grid_lat < 0 {
		grid_lat = 0
	}
	return grid_lat
}

func WrapGridLng(grid_lng, lng_tiles int) int {
	for grid_lng >= lng_tiles {
		grid_lng -= lng_tiles
	}
	for grid_lng < 0 {
		grid_lng += lng_tiles
	}
	return grid_lng
}
