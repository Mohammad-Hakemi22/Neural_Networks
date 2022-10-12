package main

import (
	"fmt"

	ts "gorgonia.org/tensor"
)

func main() {
	rawInput := []float32{
		1.0, 2.0, 3.0, 2.5,
		2.0, 5.0, -1.0, 2.0,
		-1.5, 2.7, 3.3, -0.8,
	}
	rawWeight := []float32{
		0.2, 0.8, -0.5, 1.0,
		0.5, -0.91, 0.26, 0.5,
		-0.26, -0.27, 0.17, 0.87,
	}
	rawBias := makeBias(3, []float32{2.0, 3.0, 0.5})
	input := ts.New(ts.WithBacking(rawInput), ts.WithShape(3, 4))
	weight := ts.New(ts.WithBacking(rawWeight), ts.WithShape(3, 4))
	bias := ts.New(ts.WithBacking(rawBias), ts.WithShape(3, 3))
	weight.T()
	output, err := ts.Dot(input, weight)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
	out, err := ts.Add(output, bias)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)
}

func makeBias(neuronsNum int, bias []float32) []float32 {
	biasLen := len(bias)
	res := make([]float32, neuronsNum*biasLen)
	for i := 0; i < neuronsNum; i++ {
		for j := 0; j < biasLen; j++ {
			res[j] = bias[j]
		}
	}
	return res
}
