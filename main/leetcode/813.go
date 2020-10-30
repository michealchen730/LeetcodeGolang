package main

import (
	"math"
)

func largestSumOfAverages(A []int, K int) float64 {
	//累加数组
	B := make([]int, len(A)+1)
	sum := 0
	for k, v := range A {
		sum += v
		B[k+1] = sum
	}

	dp := make([][]float64, K+1)
	for i, _ := range dp {
		dp[i] = make([]float64, len(A)+1)
	}

	//边界值初始化
	for j := 1; j <= len(A); j++ {
		dp[1][j] = float64(B[j]) / float64(j)
	}

	for i := 2; i <= K; i++ {
		for j := i; j <= len(A); j++ {
			for n := j; n >= i; n-- {
				dp[i][j] = math.Max(dp[i][j], dp[i-1][n-1]+float64(B[j]-B[n-1])/float64(j-n+1))
			}
		}
	}

	return dp[K][len(A)]
}
