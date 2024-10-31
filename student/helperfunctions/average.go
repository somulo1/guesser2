package helperfunctions

func Average(num []float64) float64 {
	var sum float64
	for _, numbers := range num {
		sum = sum + numbers
	}
	result := sum / float64(len(num))
	return result
}
