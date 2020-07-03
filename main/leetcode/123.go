package main

func maxProfit123(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dp1 := make([]int, len(prices)+1)
	mx := prices[0]
	res := 0
	for i := 1; i < len(dp1); i++ {
		mx = min(mx, prices[i-1])
		if prices[i-1] > mx {
			dp1[i] = prices[i-1] - mx
		}
		dp1[i] = max(dp1[i], dp1[i-1])
	}
	for i := 2; i < len(dp1); i++ {
		for j := i - 1; j > 0; j-- {
			if prices[j-1] < prices[i-1] {
				res = max(res, prices[i-1]-prices[j-1]+dp1[j-1])
			}
		}
	}
	return res
}
