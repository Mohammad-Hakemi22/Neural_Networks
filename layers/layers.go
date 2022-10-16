package layers

import (
	"github.com/Mohammad-Hakemi22/NN/utility"
	ts "gorgonia.org/tensor"
)

// type Layer struct {
// 	Hidden []HiddenLayer
// 	Input  InputLayer
// 	Dense  DenseLayer
// }

type DenseLayer struct {
	NeuronsNumber int
	InputNumber   int
	Weight        ts.Tensor
	Bias          ts.Tensor
	Output        ts.Tensor
}

type InputLayer struct {
	InputShape ts.Shape
	Input      ts.Tensor
}

func (d *DenseLayer) MakeBias(sh ts.Shape) {
	d.Bias, _ = ts.Repeat(d.Bias, 0, sh[0])
}

func (d *DenseLayer) Initialize() {
	weight := ts.New(ts.WithBacking(ts.Random(ts.Float64, d.InputNumber*d.NeuronsNumber)), ts.WithShape(d.InputNumber, d.NeuronsNumber))
	weight, err := weight.MulScalar(0.01, false)
	utility.ErrorHandler("something went wrong in multiply scaler to dense", err)
	d.Weight = weight
	d.Bias = ts.New(ts.WithBacking(make([]float64, d.NeuronsNumber)), ts.WithShape(1, d.NeuronsNumber))
}

func (d *DenseLayer) Forward(input ts.Tensor) {
	dp, err := ts.Dot(input, d.Weight)
	utility.ErrorHandler("something went wrong in dot product", err)
	d.MakeBias(input.Shape())
	out, err := ts.Add(dp, d.Bias)
	utility.ErrorHandler("something went wrong in add", err)
	d.Output = out
}
