package main

import (
	"math"
)

func concatenatedBinary(n int) int {
	m := 1000000007
	i := 2
	res := 1
	for n >= i {
		res *= int(math.Pow(2, float64(lenBinary(i))))
		res += i
		res %= m
		i++
	}
	return res
}
