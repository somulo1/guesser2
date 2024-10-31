package helperfunctions

func Median(num []float64) float64 {
	length := len(num)
	var median float64
	if length%2 == 0 {
		median = ((num[length/2]) + (num[length/2-1])) / 2
	} else {
		median = num[length/2]
	}
	return median
}
