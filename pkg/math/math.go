package math

import "math"

func PowInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
