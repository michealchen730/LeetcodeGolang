package main

func maximumWealth(accounts [][]int) int {
	res := 0
	for _, v := range accounts {
		w := 0
		for _, v2 := range v {
			w += v2
		}
		res = max(res, w)
	}
	return res
}
