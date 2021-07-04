package main

//这能转化成背包问题，我是想不到的
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}
	vol := sum / 2
	dp := make([]bool, vol+1)
	dp[0] = true
	for _, v := range stones {
		for j := vol; j >= 0; j-- {
			if j >= v {
				dp[j] = dp[j-v] || dp[j]
			}
		}
	}
	for i := vol; i >= 0; i-- {
		if dp[i] {
			return sum - 2*i
		}
	}
	return 0
}
