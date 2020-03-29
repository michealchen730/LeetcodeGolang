package main

import (
	"math"
	"strconv"
)

func findNthDigit(n int) int {
	if n <= 9 {
		return n
	}
	d, length, pre := 0, 0, 0
	for n > length {
		pre = length
		length += int(math.Pow(10, float64(d))) * 9 * (d + 1)
		d++
	}
	temp := (n-pre-1)/d + int(math.Pow(10, float64(d-1)))
	str, dig := strconv.Itoa(temp), (n-pre)%d
	if dig == 0 {
		dig = d
	}
	res, _ := strconv.Atoi(str[dig-1 : dig])
	return res
}
