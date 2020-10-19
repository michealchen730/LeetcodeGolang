package main

import (
	"math"
)

func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	sum := 1
	for j := 2; j <= int(math.Sqrt(float64(num))); j++ {
		if num%j == 0 {
			sum += j
			if num/j != j {
				sum += num / j
			}
		}
	}
	if sum == num {
		return true
	}
	return false
}
