package main

func translateNum(num int) int {
	if num == 0 {
		return 1
	}
	var nums []int
	for num != 0 {
		nums = append([]int{num % 10}, nums...)
		num /= 10
	}
	dp := make([]int, len(nums)+1)
	dp[0], dp[1] = 1, 1
	for i := 1; i < len(nums); i++ {
		dp[i+1] += dp[i]
		if (nums[i] <= 5 && (nums[i-1] == 2 || nums[i-1] == 1)) || (nums[i] > 5 && nums[i-1] == 1) {
			dp[i+1] += dp[i-1]
		}
	}
	return dp[len(dp)-1]
}
