package main

func numEnclaves(A [][]int) int {
	res := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			if A[i][j] == 1 {
				stack := [][]int{[]int{i, j}}
				A[i][j] = 0
				flag, num := true, 0
				//bfs
				for len(stack) != 0 {
					temp := stack[0]
					num++
					a, b, c, d := isOut(A, temp[0]+1, temp[1]), isOut(A, temp[0]-1, temp[1]), isOut(A, temp[0], temp[1]+1), isOut(A, temp[0], temp[1]-1)
					if a == 2 || b == 2 || c == 2 || d == 2 {
						flag = false
					}
					if a == 1 {
						A[temp[0]+1][temp[1]] = 0
						stack = append(stack, []int{temp[0] + 1, temp[1]})
					}
					if b == 1 {
						A[temp[0]-1][temp[1]] = 0
						stack = append(stack, []int{temp[0] - 1, temp[1]})
					}
					if c == 1 {
						A[temp[0]][temp[1]+1] = 0
						stack = append(stack, []int{temp[0], temp[1] + 1})
					}
					if d == 1 {
						A[temp[0]][temp[1]-1] = 0
						stack = append(stack, []int{temp[0], temp[1] - 1})
					}
					stack = stack[1:]
				}
				if flag {
					res += num
				}
			}
		}
	}
	return res
}

func isOut(A [][]int, i, j int) int {
	if i < 0 || i >= len(A) || j < 0 || j >= len(A[0]) {
		return 2
	} else {
		return A[i][j]
	}
}
