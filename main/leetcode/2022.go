package main

func construct2DArray(original []int, m int, n int) [][]int {
	var res [][]int
	if m*n != len(original) {
		return res
	}
	for t := n; t <= len(original); t += n {
		res = append(res, original[t-n:n])
	}
	return res
}
