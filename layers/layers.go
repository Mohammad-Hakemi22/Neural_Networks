package layers

import (
	ts "gorgonia.org/tensor"
)

type Layer struct {
	Hidden []HiddenLayer
	Input  InputLayer
	Dense  DenseLayer
}

type HiddenLayer struct {
	NeuronsNumber int64
	Weight        ts.Tensor
	Bias          ts.Tensor
}

type InputLayer struct {
	InputShape ts.Shape
	Input      ts.Tensor
}

type DenseLayer struct {
}

func (l *HiddenLayer) MakeBias() {
	l.Bias, _ = ts.Repeat(l.Bias, 0, int(l.NeuronsNumber))
}
