package main

func findJudge(N int, trust [][]int) int {
	trs := make([]int, N)
	usd := make([]int, N)
	for _, v := range trust {
		usd[v[0]] = 1
		trs[v[1]]++
	}
	for k, v := range trs {
		if v == N-1 {
			if usd[k] == 0 {
				return k + 1
			} else {
				return -1
			}
		}
	}
	return -1
}
