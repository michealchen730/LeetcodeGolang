package main

func lexicalOrder(n int) []int {
	var res []int
	for i := 1; i <= 9; i++ {
		dfs386(&res, i, n)
	}
	return res
}

func dfs386(arr *[]int, i, n int) {
	if i > n {
		return
	} else {
		*arr = append(*arr, i)
		i *= 10
		for j := 0; j <= 9; j++ {
			dfs386(arr, i+j, n)
		}
	}
}
