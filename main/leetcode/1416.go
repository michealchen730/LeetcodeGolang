package main

import (
	"math"
	"strconv"
)

func numberOfArrays(s string, k int) int {
	test := k
	bits := 0
	length := len(s)
	for test != 0 {
		test = test / 10
		bits++
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i < len(dp); i++ {
		sum := 0
		ele, _ := strconv.Atoi(s[length-i : length-i+1])
		if ele != 0 {
			for j := 1; j <= bits; j++ {
				if length-i+j <= len(s) {
					temp, _ := strconv.Atoi(s[length-i : length-i+j])
					if temp >= 1 && temp <= k {
						sum += dp[i-j] * 1
					}
				}
			}
		}
		dp[i] = sum % (int(math.Pow(10, 9)) + 7)
	}
	return dp[length]
}
