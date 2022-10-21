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
	// layer1, _ := arch.BuildDenseLayer(3, []float32{
	// 	0.2, 0.8, -0.5, 1.0,
	// 	0.5, -0.91, 0.26, -0.5,
	// 	-0.26, -0.27, 0.17, 0.87,
	// }, []float32{
	// 	2.0, 3.0, 0.5,
	// }, []int{3, 4}, []int{1, 3})

	// layer2, _ := arch.BuildDenseLayer(3, []float32{
	// 	0.1, -0.14, 0.5,
	// 	-0.5, 0.12, -0.33,
	// 	-0.44, 0.73, -0.13,
	// }, []float32{
	// 	-1.0, 2.0, -0.5,
	// }, []int{3, 3}, []int{1, 3})

	// input, _ := arch.BuildInputLayer([]float32{
	// 	1.0, 2.0, 3.0, 2.5,
	// 	2.0, 5.0, -1.0, 2.0,
	// 	-1.5, 2.7, 3.3, -0.8,
	// }, []int{3, 4})

	// layer1.Weight.T()
	// layer1_output, err := ts.Dot(input.Input, layer1.Weight)
	// utility.ErrorHandler("something went wrong in dot product", err)
	// layer1.MakeBias()
	// out, err := ts.Add(layer1_output, layer1.Bias)
	// utility.ErrorHandler("something went wrong in adding", err)
	// layer2.Weight.T()
	// layer2_output, err := ts.Dot(out, layer2.Weight)
	// utility.ErrorHandler("something went wrong in dot product", err)
	// layer2.MakeBias()
	// out_final, err := ts.Add(layer2_output, layer2.Bias)
	// utility.ErrorHandler("something went wrong in adding", err)
	// fmt.Println(out_final)
	// plot.Scatter()
	input_data, _ := datasets.SpiralData1(100, 3)
	l1 := layers.DenseLayer{InputNumber: 2, NeuronsNumber: 4}
	l1.Initialize()
	fmt.Println("-------------Weight------------")
	fmt.Println(l1.Weight)
	fmt.Println("--------------Bias-----------")
	fmt.Println(l1.Bias)
	fmt.Println("-------------output------------")
	ReLU := activationfunc.NewReLU()
	input := ts.New(ts.WithBacking(input_data), ts.WithShape(50, 2))
	l1.Forward(input)
	ReLU.Forward(l1.Output)
	// fmt.Println(l1.Output)
	fmt.Println(ReLU.Output)
}
