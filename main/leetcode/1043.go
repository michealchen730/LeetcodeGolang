package main

func maxSumAfterPartitioning(A []int, K int) int {
	dp := make([]int, len(A))
	for i := 0; i < len(dp); i++ {
		mx := A[i]
		dp[i] = mx
		for j := 0; j < K; j++ {
			if i-j >= 0 {
				mx = max(mx, A[i-j])
				t := mx * (j + 1)
				if i-j-1 >= 0 {
					t += dp[i-j-1]
				}
				dp[i] = max(dp[i], t)
			}
		}
	}
	return dp[len(dp)-1]
}
