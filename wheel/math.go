package wheel

import "math"

// Phi represents the Golden Ratio.
func Phi() float64 {
	return (1 + math.Sqrt(5)) / 2
}

// MaxInt returns the larger of x or y.
func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// MinInt returns the smaller of x or y.
func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
