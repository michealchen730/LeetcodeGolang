package main

func matrixReshape(nums [][]int, r int, c int) [][]int {
	row, col := len(nums), len(nums[0])
	if r*c != row*col {
		return nums
	}
	tmp, last := 0, row*col
	res := make([][]int, r)
	for k, _ := range res {
		res[k] = make([]int, c)
	}
	for tmp < last {
		r1, c1 := tmp/col, tmp%col
		r2, c2 := tmp/c, tmp%c
		res[r2][c2] = nums[r1][c1]
		tmp++
	}
	return res
}
