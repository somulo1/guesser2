package helperfunctions

import "math"

func Standard(num []float64) float64 {
	variance := Variance(num)
	return math.Sqrt(variance)
}
