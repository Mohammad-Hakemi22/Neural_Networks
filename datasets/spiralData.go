package datasets

import (
	"math/rand"
	"time"
)

type coordinate struct {
	x float64
	y float64
}

func SpiralData(samples, classes int) ([]coordinate, []int) {
	coor := make([]coordinate, 0, samples)
	class := make([]int, 0, samples)
	classNum := int(samples / classes)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < samples; i++ {
		coor = append(coor, coordinate{x: rand.Float64() - rand.Float64(), y: rand.Float64() - rand.Float64()})
	}
	for c := 0; c < classes; c++ {
		for j := 0; j < classNum; j++ {
			class = append(class, c)
		}
	}
	return coor, class
}
