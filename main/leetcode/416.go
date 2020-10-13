package main

import (
	"sort"
)

//这题另外有一种动态规划的解法
//此外，请参考698题，相当于是这题的升级版
//于是，这里无法通过测试用例的原因是，这题不允许穷举，所以把门槛设置的特别高
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	//如果所有元素和为奇数，那么必然无法分配
	if sum%2 != 0 {
		return false
	}
	target := sum / 2

	sort.Ints(nums)
	row := len(nums) - 1
	if nums[row] > target {
		return false
	}

	dp := make([][]bool, len(nums))
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = true
	}
	//dp[i][j] 表示从数组的 [0,i] 下标范围内选取若干个正整数（可以是 0 个），是否存在一种选取方案使得被选取的正整数的和等于 j
	dp[0][nums[0]] = true

	for i := 1; i < len(dp); i++ {
		v := nums[i]
		for j := 1; j <= target; j++ {
			if j >= v {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(dp)-1][target]
}
