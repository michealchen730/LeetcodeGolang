package main

func maxProfit714(prices []int, fee int) int {
	cash, hold := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		cash := max(cash, prices[i]+hold-fee)
		hold = max(hold, cash-prices[i])
	}
	return max(cash, hold)
}
