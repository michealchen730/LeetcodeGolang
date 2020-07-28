package main

func countSubstrings(s string) int {
	res := 1
	dp := make([][]bool, len(s))
	for k, _ := range dp {
		dp[k] = make([]bool, len(s))
	}
	dp[0][0] = true
	for i := 1; i < len(dp); i++ {
		dp[i][i] = true
		res++
		for j := i - 1; j >= 0; j-- {
			if s[j] == s[i] && (j+1 > i-1 || dp[j+1][i-1]) {
				dp[j][i] = true
				res++
			}
		}
	}
	return res
}
