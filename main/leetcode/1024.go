package main

import "math"

func videoStitching(clips [][]int, T int) int {
	dp := make([]int, T+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= T; i++ {
		for _, v := range clips {
			if v[0] < i && i <= v[1] && dp[v[0]]+1 < dp[i] {
				dp[i] = dp[v[0]] + 1
			}
		}
	}
	if dp[T] == math.MaxInt32 {
		return -1
	}
	return dp[T]
}
