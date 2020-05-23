package main

func stoneGame(piles []int) bool {
	dp := make([][]int, len(piles))
	for k, _ := range dp {
		dp[k] = make([]int, len(piles))
	}
	sum := make([][]int, len(piles))
	for k, _ := range sum {
		sum[k] = make([]int, len(piles))
	}
	for k := len(piles) - 1; k >= 0; k-- {
		s := 0
		for j := k; j < len(piles); j++ {
			s += piles[j]
			sum[k][j] = s
		}
	}

	for k := len(piles) - 1; k >= 0; k-- {
		for j := k; j < len(piles); j++ {
			if k == j {
				dp[k][j] = piles[k]
				continue
			}
			if j == k+1 {
				dp[k][j] = max(piles[k], piles[j])
			} else {
				dp[k][j] = max(piles[k]+sum[k+1][j]-dp[k+1][j], piles[j]+sum[k][j-1]-dp[k][j-1])
			}
		}
	}
	return dp[0][len(dp[0])-1] > sum[0][len(dp)-1]-dp[0][len(dp[0])-1]
}
