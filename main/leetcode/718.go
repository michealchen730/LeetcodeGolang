package main

func findLength(A []int, B []int) int {
	dp := make([][]int, len(A)+1)
	for k, _ := range dp {
		dp[k] = make([]int, len(B)+1)
	}
	mx := 0
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				mx = max(dp[i][j], mx)
			}
		}
	}
	return mx
}
