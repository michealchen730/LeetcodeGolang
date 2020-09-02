package main

func numDistinct(s string, t string) int {
	dp := make([][]int, len(s)+1)
	for k, _ := range dp {
		dp[k] = make([]int, len(t)+1)
	}

	for i := 0; i < len(dp); i++ {
		dp[i][0] = 1
	}

	//dp[i][j]代表s前i个字符中包含了多少种t前[j]字符的情况
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if s[i] == t[j] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(s)][len(t)]
}
