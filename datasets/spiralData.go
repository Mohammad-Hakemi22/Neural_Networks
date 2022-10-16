package datasets

import (
	"fmt"
	"math/rand"
	"time"
)


//TODO: use sin and cos for points coordinate

func SpiralData(samples, classes int) ([][]float64, []int) {
	coor := make([][]float64, 0, samples)
	class := make([]int, 0, samples)
	var classNum int = 0
	if samples % classes == 0 {
		classNum = int(samples / classes)
	} else {
		classNum = int(samples / classes) + 1
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < samples; i++ {
		coor = append(coor, []float64{rand.NormFloat64(), rand.NormFloat64()})
	}
	for c := 0; c < classes; c++ {
		for j := 0; j < classNum; j++ {
			class = append(class, c)
		}
	}
	if lenght := len(class) - samples; lenght != 0 {
		class = class[:samples]
	}
	fmt.Println(coor)
	return coor, class
}

func SpiralData1(samples, classes int) ([]float64, []int) {
	coor := []float64{}
	class := make([]int, 0, samples)
	var classNum int = 0
	if samples % classes == 0 {
		classNum = int(samples / classes)
	} else {
		classNum = int(samples / classes) + 1
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < samples; i++ {
		coor = append(coor, rand.NormFloat64())
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
