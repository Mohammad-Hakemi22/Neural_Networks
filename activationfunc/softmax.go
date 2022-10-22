package activationfunc

import (
	"github.com/Mohammad-Hakemi22/NN/utility"
	ts "gorgonia.org/tensor"
)

type SoftMax struct {
	Output ts.Tensor
}

func NewSoftMax() *SoftMax {
	return &SoftMax{Output: ts.New(ts.Of(ts.Float64))}
}

func (s *SoftMax) Forward(input ts.Tensor, axis int) {
	out, err := ts.SoftMax(input, axis)
	utility.ErrorHandler("something went wrong in calculate softmax:", err)
	s.Output = out
}

// func SoftMax(input []float64) []float64 {
// 	output := []float64{}
// 	var sum float64 = 0
// 	for _, val := range input {
// 		sum += math.Pow(math.E, val)
// 	}
// 	for _, val := range input {
// 		output = append(output, val/sum)
// 	}
// 	return output
// }
