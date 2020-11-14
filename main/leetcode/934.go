package main

func shortestBridge(A [][]int) int {
	//先用一层dfs标记完整座岛
	//再用一整个数组记录这座岛
	//最后使用bfs得出最小数目

	var island [][]int

	var dfs func(i, j int)
	dfs = func(i, j int) {
		A[i][j] = 2
		island = append(island, []int{i, j})
		if i-1 >= 0 && A[i-1][j] == 1 {
			dfs(i-1, j)
		}
		if i+1 < len(A) && A[i+1][j] == 1 {
			dfs(i+1, j)
		}
		if j-1 >= 0 && A[i][j-1] == 1 {
			dfs(i, j-1)
		}
		if j+1 < len(A[0]) && A[i][j+1] == 1 {
			dfs(i, j+1)
		}
	}
SE:
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			if A[i][j] == 1 {
				dfs(i, j)
				break SE
			}
		}
	}

	res := 0
FI:
	for len(island) != 0 {
		var queue [][]int
		for _, v := range island {
			if v[0]-1 >= 0 {
				if A[v[0]-1][v[1]] == 1 {
					break FI
				} else if A[v[0]-1][v[1]] == 0 {
					A[v[0]-1][v[1]] = 2
					queue = append(queue, []int{v[0] - 1, v[1]})
				}
			}
			if v[0]+1 < len(A) {
				if A[v[0]+1][v[1]] == 1 {
					break FI
				} else if A[v[0]+1][v[1]] == 0 {
					A[v[0]+1][v[1]] = 2
					queue = append(queue, []int{v[0] + 1, v[1]})
				}
			}
			if v[1]-1 >= 0 {
				if A[v[0]][v[1]-1] == 1 {
					break FI
				} else if A[v[0]][v[1]-1] == 0 {
					A[v[0]][v[1]-1] = 2
					queue = append(queue, []int{v[0], v[1] - 1})
				}
			}
			if v[1]+1 < len(A[0]) {
				if A[v[0]][v[1]+1] == 1 {
					break FI
				} else if A[v[0]][v[1]+1] == 0 {
					A[v[0]][v[1]+1] = 2
					queue = append(queue, []int{v[0], v[1] + 1})
				}
			}
		}
		island = queue
		res++
	}

	return res
}
