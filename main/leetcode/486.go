package main

func PredictTheWinner(nums []int) bool {
	dp := make([][]int, len(nums))
	for k, _ := range dp {
		dp[k] = make([]int, len(nums))
	}
	sum := make([][]int, len(nums))
	for k, _ := range sum {
		sum[k] = make([]int, len(nums))
	}
	for k := len(nums) - 1; k >= 0; k-- {
		s := 0
		for j := k; j < len(nums); j++ {
			s += nums[j]
			sum[k][j] = s
		}
	}

	for k := len(nums) - 1; k >= 0; k-- {
		for j := k; j < len(nums); j++ {
			if k == j {
				dp[k][j] = nums[k]
				continue
			}
			if j == k+1 {
				dp[k][j] = max(nums[k], nums[j])
			} else {
				dp[k][j] = max(nums[k]+sum[k+1][j]-dp[k+1][j], nums[j]+sum[k][j-1]-dp[k][j-1])
			}
		}
	}
	return dp[0][len(dp[0])-1] >= sum[0][len(dp)-1]-dp[0][len(dp[0])-1]
}
