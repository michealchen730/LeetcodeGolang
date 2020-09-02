package main

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	var res [][]int
	res = append(res, []int{1})
	if numRows == 1 {
		return res
	}
	res = append(res, []int{1, 1})
	t := 2
	for t < numRows {
		arr := []int{1}
		for i := 1; i < t; i++ {
			arr = append(arr, res[len(res)-1][i-1]+res[len(res)-1][i])
		}
		arr = append(arr, 1)
		t++
		res = append(res, arr)
	}
	return res
}
