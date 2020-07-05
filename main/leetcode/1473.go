package main

import (
	"math"
)

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	dp := make([][][]int, m)
	//dp[i][j][k]表示前i个房子涂成j个街区，且第i个房子是颜色k的最小花费
	for k, _ := range dp {
		dp[k] = make([][]int, m+1)
		for i, _ := range dp[k] {
			dp[k][i] = make([]int, n)
			for d, _ := range dp[k][i] {
				dp[k][i][d] = math.MaxInt32
			}
		}
	}
	//初始化第一个房子的涂色方案，先设置为最大值（无效值）
	//如果第一个房子有颜色，那么仅更新这个颜色
	if houses[0] != 0 {
		dp[0][1][houses[0]-1] = 0
	} else {
		//否则更新所有涂色方案
		for j := 0; j < n; j++ {
			dp[0][1][j] = cost[0][j]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j <= i+1; j++ {
			if houses[i] != 0 {
				dp[i][j][houses[i]-1] = dp[i-1][j][houses[i]-1]
				for k := 0; k < n; k++ {
					if k != houses[i]-1 {
						dp[i][j][houses[i]-1] = min(dp[i][j][houses[i]-1], dp[i-1][j-1][k])
					}
				}
			} else {
				for k := 0; k < n; k++ {
					if dp[i-1][j][k] != math.MaxInt32 {
						dp[i][j][k] = dp[i-1][j][k] + cost[i][k]
					}
					for t := 0; t < n; t++ {
						if t != k {
							dp[i][j][k] = min(dp[i][j][k], dp[i-1][j-1][t]+cost[i][k])
						}
					}
				}
			}
		}
	}
	res := dp[m-1][target][0]
	for i := 0; i < n; i++ {
		res = min(dp[m-1][target][i], res)
	}

	if res == math.MaxInt32 {
		return -1
	}
	return res
}
