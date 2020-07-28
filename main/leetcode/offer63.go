package main

func maxProfitO63(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	mn := prices[0]
	res := 0
	for _, v := range prices {
		if v > mn {
			res = max(res, v-mn)
		} else {
			mn = v
		}
	}
	return res
}
