package main

func closedIsland(grid [][]int) int {
	//判断封闭岛屿的条件
	//BFS，起点为某个为0的坐标，走遍整个岛屿

	res := 0

	//初始化一个矩阵，表示某个坐标是否被走过了
	row := len(grid)
	col := len(grid[0])

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 0 {
				grid[i][j] = 2
				queue := [][]int{[]int{i, j}}
				flag := 0 //flag用来表示岛屿中是否出现了位于边界上的点，有的话，flag置1
				//开始BFS
				for len(queue) != 0 {
					temp := queue[0]
					if temp[0] == 0 || temp[0] == row-1 || temp[1] == 0 || temp[1] == col-1 {
						flag = 1
					}
					if checkIsland(grid, temp[0]+1, temp[1]) == 0 {
						queue = append(queue, []int{temp[0] + 1, temp[1]})
						grid[temp[0]+1][temp[1]] = 2
					}
					if checkIsland(grid, temp[0]-1, temp[1]) == 0 {
						queue = append(queue, []int{temp[0] - 1, temp[1]})
						grid[temp[0]-1][temp[1]] = 2
					}
					if checkIsland(grid, temp[0], temp[1]+1) == 0 {
						queue = append(queue, []int{temp[0], temp[1] + 1})
						grid[temp[0]][temp[1]+1] = 2
					}
					if checkIsland(grid, temp[0], temp[1]-1) == 0 {
						queue = append(queue, []int{temp[0], temp[1] - 1})
						grid[temp[0]][temp[1]-1] = 2
					}
					queue = queue[1:]
				}
				if flag == 0 {
					res++
				}
			}
		}
	}

	return res

}
