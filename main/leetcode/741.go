package main

import (
	"math"
)

func cherryPickup741(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][][]int, 2*m)
	for k, _ := range dp {
		dp[k] = make([][]int, n)
		for l, _ := range dp[k] {
			dp[k][l] = make([]int, n)
			for o, _ := range dp[k][l] {
				dp[k][l][o] = math.MinInt32
			}
		}
	}
	dp[0][0][0] = grid[0][0]

	//可以视为两个玩家在行走
	//dp[i][j][k]表示第i步，第一人在第j列，第二人在第k列的情况
	for i := 1; i <= 2*m-2; i++ {
	FI:
		for j := 0; j < n; j++ {
		SC:
			for k := 0; k < n; k++ {
				if j > i {
					break FI
				}
				if k > i {
					break SC
				}

				//首先，排除障碍的情况
				if i-j >= len(grid) || i-k >= len(grid) || grid[i-j][j] == -1 || grid[i-k][k] == -1 {
					continue
				}
				//不是障碍
				temp := grid[i-j][j]
				if j != k {
					temp += grid[i-k][k]
				}
				if i-1 >= j && i-1 >= k {
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j][k])
				}
				if j-1 >= 0 && i-1 >= k {
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j-1][k])
				}
				if i-1 >= j && k-1 >= 0 {
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j][k-1])
				}
				if j-1 >= 0 && k-1 >= 0 {
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j-1][k-1])
				}
				if dp[i][j][k] != math.MinInt32 {
					dp[i][j][k] += temp
				}
			}
		}
	}
	if dp[2*m-2][m-1][m-1] == math.MinInt32 {
		return 0
	}
	return dp[2*m-2][m-1][m-1]
}
