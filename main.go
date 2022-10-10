package main

import (
	"errors"
	"fmt"
)

type Neuron struct {
	Bias   float32
	Weight []float32
}

func main() {
	input := []float32{1.0, 2.0, 3.0, 2.5}
	neurons := []Neuron{
		{Bias: 2, Weight: []float32{0.2, 0.8, -0.5, 1}},
		{Bias: 3, Weight: []float32{0.5, -0.91, 0.26, -0.5}},
		{Bias: 0.5, Weight: []float32{-0.26, -0.27, 0.17, 0.87}},
	}
	fmt.Println(sumation(input, neurons))
}

func sumation(input []float32, neurons []Neuron) ([]float32, error) {
	output := make([]float32, len(neurons))
	for i := 0; i < len(neurons); i++ {
		dot_product, _ := dot(input, neurons[i].Weight)
		output[i] = dot_product + neurons[i].Bias
	}
	return output, nil
}

func dot(a, b []float32) (float32, error) {
	if len(a) != len(b) {
			return 0, errors.New("lenght of input and weight must be same")
		}
	var output float32 = 0.0
	for i := 0; i < len(a); i++ {
		output += a[i] * b[i]
	}
	return output, nil
}