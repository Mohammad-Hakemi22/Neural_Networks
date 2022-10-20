package activationfunc

import "math"

func SoftMax(input []float64) []float64 {
	output := []float64{}
	var sum float64 = 0
	for _, val := range input {
		sum += math.Pow(math.E, val)
	}
	for _, val := range input {
		output = append(output, val/sum)
	}
	return output
}