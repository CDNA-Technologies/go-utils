package number

import "math"

func IsWhole(num float64) bool {
	if math.IsInf(num, 0) || math.IsNaN(num) {
		return false
	}
	return math.Trunc(num) == num
}
