package game

import "math"

func vectorLen(x, y float64) float64 {
	var len = math.Sqrt(x*x + y*y)
	return len
}
func vectorNormalize(x, y float64) (float64, float64) {
	len := vectorLen(x, y)
	if len == 0 {
		return 0.01, 0.01
	}
	return x / len, y / len
}
func vectorAdd(x, y, d, min, max float64) (float64, float64) {
	len := vectorLen(x, y)
	nx, ny := vectorNormalize(x, y)
	len2 := len + d
	if len2 < min {
		len2 = min
	}
	if len2 > max {
		len2 = max
	}

	x2 := nx * len2
	y2 := ny * len2

	return x2, y2
}
func clamp255(i int) uint8 {
	if i < 0 {
		return 0
	}
	if i > 255 {
		return 255
	}
	return uint8(i)
}
func angleNormal(a float64) float64 {
	for true {
		if a < 0 {
			a = a + 360
		} else if a > 360 {
			a = a - 360
		} else {
			break
		}
	}
	return a
}
