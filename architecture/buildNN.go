package architecture

import (
	layer "github.com/Mohammad-Hakemi22/NN/layers"
	ts "gorgonia.org/tensor"
)

func BuildDenseLayer(neuronsNum int, weight, bias []float32, weightShape, biasShape []int) (*layer.DenseLayer, error) {
	hl := layer.DenseLayer{
		NeuronsNumber: neuronsNum,
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
