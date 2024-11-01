package helperfunctions

// import "fmt"
func LinearRg(x, y []float64) (float64, float64) {
	// y = mx + c
	var numerator float64
	var denominator float64
	meany := Average(y)
	meanx := Average(x)

	for i := 0; i < len(x); i++ {
		if i != len(x)-1{
			numerator += (x[i] - meanx) * (y[i] - meany)
			denominator += (x[i] - meanx) * (x[i] - meanx)
		}
	
	}
	m := numerator / denominator
	a := meany - m*meanx
	return m, a
}
