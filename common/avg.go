package common

//AvgFloat64 calculates avg value for slice of float64
func AvgFloat64(in []float64) float64 {
	var sum float64

	for _, value := range in {
		sum += value
	}

	return sum / float64(len(in))
}
