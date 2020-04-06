package main

import (
	"math"
	"strconv"
)

func maximum69Number(num int) int {
	s := strconv.Itoa(num)
	for i := 0; i < len(s); i++ {
		if s[i] == '6' {
			return num + 3*int(math.Pow(10, float64(len(s)-i-1)))
		}
	}
	return num
}
