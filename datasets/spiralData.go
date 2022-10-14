package datasets

import (
	"math/rand"
	"time"
)

type coordinate struct {
	X float64
	Y float64
}

//TODO: use sin and cos for points coordinate

func SpiralData(samples, classes int) ([]coordinate, []int) {
	coor := make([]coordinate, 0, samples)
	class := make([]int, 0, samples)
	var classNum int = 0
	if samples % classes == 0 {
		classNum = int(samples / classes)
	} else {
		classNum = int(samples / classes) + 1
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < samples; i++ {
		coor = append(coor, coordinate{X: rand.Float64() - rand.Float64(), Y: rand.Float64() - rand.Float64()})
	}
	for c := 0; c < classes; c++ {
		for j := 0; j < classNum; j++ {
			class = append(class, c)
		}
	}
	if lenght := len(class) - samples; lenght != 0 {
		class = class[:samples]
	}
	return coor, class
}
