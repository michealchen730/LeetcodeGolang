package main

import (
	"math"
)

func maxProfit188(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	buy := make([][]int, len(prices))
	sell := make([][]int, len(prices))
	for m, _ := range buy {
		buy[m] = make([]int, k+1)
		sell[m] = make([]int, k+1)
	}

	for i := 1; i <= k; i++ {
		buy[0][i] = math.MinInt32
		sell[0][i] = math.MinInt32
	}
	buy[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		buy[i][0] = max(buy[i-1][0], sell[i-1][0]-prices[i])
		for j := 1; j <= k; j++ {
			//如果是i天买进，那就是sell[i-1][j]-prices[i]
			//如果不是第i天买进，那就是buy[i-1][j]
			buy[i][j] = max(buy[i-1][j], sell[i-1][j]-prices[i])
			//如果是i天卖出，那就是buy[i-1][j-1]+prices[i]
			//如果不是第i天卖出，那就是sell[i-1][j]
			sell[i][j] = max(sell[i-1][j], buy[i-1][j-1]+prices[i])

		}
	}
	res := 0
	for i := 1; i <= k; i++ {
		res = max(res, sell[len(prices)-1][i])
	}
	return res
}
