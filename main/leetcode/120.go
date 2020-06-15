package main

//这题除了dp还可以使用深度优先，广度优先存疑

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < i; j++ {
			temp := triangle[i-1][j] + triangle[i][j]
			if j-1 >= 0 {
				temp = min(temp, triangle[i-1][j-1]+triangle[i][j])
			}
			triangle[i][j] = temp
		}
		triangle[i][i] += triangle[i-1][i-1]
	}
	temp := triangle[len(triangle)-1][0]
	for _, v := range triangle[len(triangle)-1] {
		if v < temp {
			temp = v
		}
	}
	return temp
}
