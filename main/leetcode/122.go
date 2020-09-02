package main

func maxProfit122(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	res := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}
