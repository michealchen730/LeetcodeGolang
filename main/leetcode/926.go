package main

func minFlipsMonoIncr(S string) int {
	dp := make([][]int, len(S)+1)
	for k, _ := range dp {
		dp[k] = make([]int, 2)
	}
	for i := 1; i <= len(S); i++ {
		if S[i-1] == '1' {
			dp[i][0] = dp[i-1][0] + 1
			dp[i][1] = min(dp[i-1][0], dp[i-1][1])
		} else {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = min(dp[i-1][1], dp[i-1][0]) + 1
		}
	}
	return min(dp[len(S)][0], dp[len(S)][1])
}

//可以简化到两个变量
func minFlipsMonoIncr2(S string) int {
	i, j := 0, 0
	for k := 1; k <= len(S); k++ {
		if S[k-1] == '1' {
			i, j = i+1, min(i, j)
		} else {
			j = min(i, j) + 1
		}
	}
	return min(i, j)
}
