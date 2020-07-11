package main

func maxProfit309(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	dp := make([][3]int, n)
	//第0列表示当天持有股票的最大收益
	dp[0][0] = -prices[0]
	//第1列表示当天卖出股票的最大收益
	dp[0][1] = 0
	//第1列表示当天不持有股票，且不处于冷冻期的最大收益
	dp[0][2] = 0
	for i := 1; i < len(dp); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[n-1][0], max(dp[n-1][1], dp[n-1][2]))
}
