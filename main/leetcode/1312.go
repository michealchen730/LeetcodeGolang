package main

func minInsertions(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for k, _ := range dp {
		dp[k] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] {
				if i == j+1 {
					dp[i][j] = 0
				} else {
					dp[i][j] = dp[i-1][j+1]
				}
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j+1]) + 1
			}
		}
	}
	return dp[n-1][0]
}
