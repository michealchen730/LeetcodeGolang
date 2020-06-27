package main

func subsets(nums []int) [][]int {
	var res [][]int
	var path []int
	dfs0804(0, &res, &path, nums)
	return res
}

func dfs0804(n int, res *[][]int, path *[]int, nums []int) {
	if n == len(nums) {
		cpy := make([]int, len(*path))
		copy(cpy, *path)
		*res = append(*res, cpy)
		return
	}
	dfs0804(n+1, res, path, nums)
	*path = append(*path, nums[n])
	dfs0804(n+1, res, path, nums)
	*path = (*path)[:len(*path)-1]
}
