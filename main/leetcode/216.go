package main

func combinationSum3(k int, n int) [][]int {
	if n > 45 {
		return nil
	}
	var res [][]int
	var path []int
	for i := 1; i <= 9; i++ {
		dfs216(&res, &path, i, n, k)
	}
	return res
}

func dfs216(res *[][]int, path *[]int, i int, left int, length int) {
	*path = append(*path, i)
	left = left - i
	if len(*path) == length {
		if left == 0 {
			cpy := make([]int, len(*path))
			copy(cpy, *path)
			*res = append(*res, cpy)
		}
		*path = (*path)[:len(*path)-1]
		return
	}
	//剪枝一,比如n=23，此时path为[9]，剪除这种情况；剪枝二，当left<0时剪枝
	if length-len(*path) > 9-(*path)[len(*path)-1] || left < 0 {
		*path = (*path)[:len(*path)-1]
		return
	}
	temp := (*path)[len(*path)-1]
	for j := temp + 1; j <= 9; j++ {
		dfs216(res, path, j, left, length)
	}
	*path = (*path)[:len(*path)-1]
}
