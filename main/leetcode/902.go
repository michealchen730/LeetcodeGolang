package main

import (
	"math"
	"strconv"
)

//数位DP问题
//dp[i]表示小于等于N中最后|N|-i的合法数的个数
func atMostNGivenDigitSet(digits []string, n int) int {
	s := strconv.Itoa(n)
	K := len(s)
	dp := make([]int, K+1)
	dp[K] = 1

	for i := K - 1; i >= 0; i-- {
		si := s[i] - '0'
		for _, v := range digits {
			if v[0]-'0' < si {
				dp[i] += int(math.Pow(float64(len(digits)), float64(K-i-1)))
			} else if v[0]-'0' == si {
				dp[i] += dp[i+1]
			}
		}
	}

	//对于2345，那么只要是3位数的数字，必然小于2345，因此可以用digits里的数字任意组合
	for i := 1; i < K; i++ {
		dp[0] += int(math.Pow(float64(len(digits)), float64(i)))
	}

	return dp[0]
}
