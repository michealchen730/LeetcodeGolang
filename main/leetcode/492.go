package main

import (
	"math"
)

func constructRectangle(area int) []int {
	for i := Sqrt492(area); i >= 1; i-- {
		if area%i == 0 {
			return []int{area / i, i}
		}
	}
	return []int{0, 0}
}

func Sqrt492(x int) int {
	if x == 0 {
		return 0
	}
	C, x0 := float64(x), float64(x)
	for {
		xi := 0.5 * (x0 + C/x0)
		if math.Abs(x0-xi) < 1e-7 {
			break
		}
		x0 = xi
	}
	return int(x0)
}
