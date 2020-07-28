package main

func minimumDeleteSum(s1 string, s2 string) int {
	sum := 0
	for _, v := range s1 {
		sum += int(v)
	}
	for _, v := range s2 {
		sum += int(v)
	}
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for k, _ := range dp {
		dp[k] = make([]int, n+1)
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + int(s1[i-1])
			}
			dp[i][j] = max(dp[i][j], dp[i][j-1])
			dp[i][j] = max(dp[i][j], dp[i-1][j])
		}
	}
	return sum - dp[m][n]*2
}
