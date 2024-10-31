package helperfunctions

func Variance(num []float64) float64 {
	var variance float64
	mean := Average(num)
	for _, ch := range num {
		variance += (ch - mean) * (ch - mean)
	}
	return variance / float64(len(num))
}
