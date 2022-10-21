package activationfunc

import (
	"github.com/Mohammad-Hakemi22/NN/utility"
	ts "gorgonia.org/tensor"
)

type ReLU struct {
	Output ts.Tensor
}

func NewReLU() *ReLU {
	return &ReLU{Output: ts.New(ts.Of(ts.Float64))}
}

func (r *ReLU) Forward(input ts.Tensor) {
	zeros := ts.New(ts.WithShape(input.Shape()...), ts.Of(ts.Float64))
	out, err := ts.MaxBetween(zeros, input)
	utility.ErrorHandler("something went wrong in calculate max:", err)
	r.Output = out
}
