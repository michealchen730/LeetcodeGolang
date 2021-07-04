package main

//看了一下DP解法，优化后的思路将这题变成了一个背包问题
//func findTargetSumWays(nums []int, target int) int {
//	sum:=0
//	for _,v:=range nums{
//		sum+=v
//	}
//	if sum-target<0||(sum-target)%2!=0{
//		return 0
//	}
//	neg:=(sum-target)/2
//	dp:=make([][]int,len(nums)+1)
//	for k,_:=range dp{
//		dp[k]=make([]int,neg+1)
//		dp[k][0]=1
//	}
//	//dp[i][j]表示前i件物品和为j的方案数
//	for i:=1;i<=len(nums);i++{
//		for j:=0;j<=neg;j++{
//			dp[i][j]=dp[i-1][j]
//			if j-nums[i-1]>=0{
//				dp[i][j]+=dp[i-1][j-nums[i-1]]
//			}
//		}
//	}
//	return dp[len(nums)][neg]
//}

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum-target < 0 || (sum-target)%2 != 0 {
		return 0
	}
	neg := (sum - target) / 2
	dp := make([]int, neg+1)
	dp[0] = 1
	//dp[i][j]表示前i件物品和为j的方案数
	for i := 1; i <= len(nums); i++ {
		for j := neg; j >= 0; j-- {
			if j-nums[i-1] >= 0 {
				dp[j] += dp[j-nums[i-1]]
			}
		}
	}
	return dp[neg]
}
