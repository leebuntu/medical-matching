package maps

import "math"

func haversine(srcLat, srcLon, dstLat, dstLon float64) float64 {
	const R = 6371

	srcLatRad := srcLat * math.Pi / 180
	srcLonRad := srcLon * math.Pi / 180
	dstLatRad := dstLat * math.Pi / 180
	dstLonRad := dstLon * math.Pi / 180

	deltaLat := dstLatRad - srcLatRad
	deltaLon := dstLonRad - srcLonRad

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(srcLatRad)*math.Cos(dstLatRad)*math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func IsWithinRadius(srcLat, srcLon, dstLat, dstLon, radius float64) bool {
	distance := haversine(srcLat, srcLon, dstLat, dstLon)
	return distance <= radius
}

func GetWalkingTime(startLongitude, startLatitude, endLongitude, endLatitude float64) float64 {
	distance := haversine(startLatitude, startLongitude, endLatitude, endLongitude)
	return distance / 1.39
}
