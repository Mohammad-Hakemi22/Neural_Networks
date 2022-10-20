package activationfunc

import "math"

func ReLU(input []float64) []float64 {
	output := []float64{}
	for _, val := range output {
		output = append(output, math.Max(0, val))
	}
	return output
}