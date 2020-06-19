package main

func finalPrices(prices []int) []int {
	var res []int
	for i := 0; i < len(prices)-1; i++ {
		flag := true
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				res = append(res, prices[i]-prices[j])
				flag = false
				break
			}
		}
		if flag {
			res = append(res, prices[i])
		}
	}
	res = append(res, prices[len(prices)-1])
	return res
}
