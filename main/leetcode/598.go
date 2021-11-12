package main

func maxCount(m int, n int, ops [][]int) int {
	for _, v := range ops {
		m = min(m, v[0])
		n = min(n, v[1])
	}
	return m * n
}
