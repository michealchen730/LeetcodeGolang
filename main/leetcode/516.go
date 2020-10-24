package main

func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s)+1)
	for k, _ := range dp {
		dp[k] = make([]int, len(s)+1)
	}
	for i := len(s) - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				if j == i+1 {
					dp[i][j] = 2
				} else {
					dp[i][j] = dp[i+1][j-1] + 2
				}
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}
