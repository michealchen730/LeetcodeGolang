package main

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	ans := 0
	res := make([][]int, len(matrix))
	for k, _ := range res {
		res[k] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		res[i][0] = int(matrix[i][0] - '0')

	}
	for i := 0; i < len(matrix[0]); i++ {
		res[0][i] = int(matrix[0][i] - '0')
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				continue
			}
			if res[i-1][j-1] != 0 && res[i-1][j] != 0 && res[i][j-1] != 0 {
				res[i][j] = minOfThree(res[i-1][j-1], res[i-1][j], res[i][j-1]) + 1
			} else {
				res[i][j] = 1
			}
		}
	}
	for _, v := range res {
		for _, w := range v {
			if w > ans {
				ans = w
			}
		}
	}
	return ans * ans
}

func minOfThree(i int, j int, k int) int {
	if i <= j && i <= k {
		return i
	} else if j <= k {
		return j
	} else {
		return k
	}
}
