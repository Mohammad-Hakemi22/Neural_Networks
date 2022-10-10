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
	inputLenght := len(input)
	for _, neuron := range neurons {
		if inputLenght != len(neuron.Weight) {
			return nil, errors.New("lenght of input and weight must be same")
		}
	}
	output := make([]float32, len(neurons))
	for i := 0; i < len(neurons); i++ {
		for j := 0; j < len(input); j++ {
			output[i] += input[j] * neurons[i].Weight[j]
		}
		output[i] += neurons[i].Bias
	}
	return output, nil
}
