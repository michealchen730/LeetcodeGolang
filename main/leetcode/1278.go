package main

func palindromePartition(s string, k int) int {
	dp := make([][]int, len(s)+1)
	for i, _ := range dp {
		dp[i] = make([]int, k+1)
	}
	cost := cost1278(s)
	for i := 1; i < len(dp); i++ {
		dp[i][1] = cost[0][i-1]
	}
	for i := 2; i < len(dp); i++ {
		for j := 2; j <= i; j++ {
			dp[i][j] = i
			for m := j - 1; m < i; m++ {
				dp[i][j] = min(dp[i][j], dp[m][j-1]+cost[m][i-1])
			}
		}
	}
	return dp[len(s)][k]
}

func cost1278(s string) [][]int {
	cost := make([][]int, len(s))
	for k, _ := range cost {
		cost[k] = make([]int, len(s))
	}
	for i := len(cost) - 1; i >= 0; i-- {
		for j := i; j < len(cost[0]); j++ {
			if i == j {
				cost[i][j] = 0
				continue
			}
			if j-i == 1 {
				if s[i] == s[j] {
					cost[i][j] = 0
				} else {
					cost[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					cost[i][j] = cost[i+1][j-1]
				} else {
					cost[i][j] = cost[i+1][j-1] + 1
				}
			}
		}
	}
	return cost
}
