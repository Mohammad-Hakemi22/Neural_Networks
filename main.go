package main

import (
	"fmt"

	"github.com/Mohammad-Hakemi22/NN/layers"
	"github.com/Mohammad-Hakemi22/NN/datasets"
	"github.com/Mohammad-Hakemi22/NN/activationfunc"
	// arch "github.com/Mohammad-Hakemi22/NN/architecture"
	// "github.com/Mohammad-Hakemi22/NN/plot"
	// "github.com/Mohammad-Hakemi22/NN/utility"
	ts "gorgonia.org/tensor"
)

func main() {
	// Input Random Data
	input_data, _ := datasets.SpiralData1(100, 3)
	input := ts.New(ts.WithBacking(input_data), ts.WithShape(50, 2))

	// Layer 1
	l1 := layers.DenseLayer{InputNumber: 2, NeuronsNumber: 4}
	l1.Initialize()
	ReLU := activationfunc.NewReLU()

	// Layer 2
	l2 := layers.DenseLayer{InputNumber: 4, NeuronsNumber: 2}
	l2.Initialize()
	SoftMax := activationfunc.NewSoftMax()


	l1.Forward(input)
	ReLU.Forward(l1.Output)

	l2.Forward(ReLU.Output)
	SoftMax.Forward(l2.Output, 1)
	
	fmt.Println(SoftMax.Output)
}
