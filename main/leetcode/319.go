package main

import "math"

// 牛逼的脑筋急转弯
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n) + 1/2))
}
