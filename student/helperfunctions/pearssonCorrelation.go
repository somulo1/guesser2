// calculating pearson correlation

package helperfunctions

import "math"

func PearsonCorr(x, y []float64) float64 {
	var sumx float64
	var sumy float64
	n := float64(len(x))
	var sumxy float64
	var sumx2 float64
	var sumy2 float64
	var numerator float64
	var denominator float64

	for i := 0; i < len(x); i++ {
		sumx += x[i]
		sumx2 += x[i] * x[i]
		sumy += y[i]
		sumxy += x[i] * y[i]
		sumy2 += y[i] * y[i]

	}
	numerator = (n * sumxy) - (sumx * sumy)
	denominator = math.Sqrt(((n * sumx2) - (sumx * sumx)) * ((n * sumy2) - (sumy * sumy)))

	r := numerator / denominator

	return r
}
