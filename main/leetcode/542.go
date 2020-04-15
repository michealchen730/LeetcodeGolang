package main

func updateMatrix(matrix [][]int) [][]int {
	res := make([][]int, len(matrix))
	for k, _ := range res {
		res[k] = make([]int, len(matrix[0]))
		copy(res[k], matrix[k])
	}
	var stack [][]int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] != 0 {
				stack = append(stack, []int{i, j})
			}
		}
	}
	nums := 1
	for len(stack) != 0 {
		temp := stack
		stack = nil
		var b0 [][]int
		for _, v := range temp {
			if checkBorder(matrix, v[0], v[1]) {
				b0 = append(b0, []int{v[0], v[1]})
			} else {
				stack = append(stack, []int{v[0], v[1]})
			}
		}
		for _, v := range b0 {
			matrix[v[0]][v[1]] = 0
			res[v[0]][v[1]] = nums
		}
		nums++
	}
	return res
}

func checkBorder(mat [][]int, i, j int) bool {
	if (i+1 < len(mat) && mat[i+1][j] == 0) || (i-1 >= 0 && mat[i-1][j] == 0) || (j+1 < len(mat[0]) && mat[i][j+1] == 0) || (j-1 >= 0 && mat[i][j-1] == 0) {
		return true
	}
	return false
}
