package helperfunctions

// Function handling the lowerlimit and upperlimit.
func Limit(xa, ya []float64) (float64, float64) {
	// y = mx + c
	var y float64
	if len(ya) == 0 {
		return 0, 0
	}

	std := Standard(ya)
	m, c := LinearRg(xa, ya)

	for i := 0; i < len(xa); i++ {
		y = m*(xa[i]+1) + c
	}

	lowerLimit := float64(y - (std * 3))
	upperLimit := float64(y + (std * 3))

	return lowerLimit, upperLimit
}
