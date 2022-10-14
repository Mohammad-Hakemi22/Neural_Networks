package architecture

import (
	layer "github.com/Mohammad-Hakemi22/NN/layers"
	ts "gorgonia.org/tensor"
)

func BuildHiddenLayer(neuronsNum int, weight, bias []float32, weightShape, biasShape []int) (*layer.HiddenLayer, error) {
	hl := layer.HiddenLayer{
		NeuronsNumber: int64(neuronsNum),
		Weight: ts.New(
			ts.WithShape(weightShape...),
			ts.WithBacking(weight)),
		Bias: ts.New(
			ts.WithShape(biasShape...),
			ts.WithBacking(bias)),
	}
	return &hl, nil
}

func BuildInputLayer(input []float32, inputShape []int) (*layer.InputLayer, error) {
	il := layer.InputLayer{
		InputShape: inputShape,
		Input: ts.New(
			ts.WithShape(inputShape...),
			ts.WithBacking(input),
		),
	}
	return &il, nil
}
