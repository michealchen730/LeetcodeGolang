package main

func orangesRotting(grid [][]int) int {
	var stack [][]int
	freshO, res := 0, 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				stack = append(stack, []int{i, j})
			}
			if grid[i][j] == 1 {
				freshO++
			}
		}
	}
	for len(stack) != 0 {
		nums := len(stack)
		if grid[stack[0][0]][stack[0][1]] == 3 {
			res++
		}
		for i := 0; i < nums; i++ {
			temp := stack[0]
			if checkFresh(grid, temp[0]-1, temp[1]) {
				grid[temp[0]-1][temp[1]] = 3
				freshO--
				stack = append(stack, []int{temp[0] - 1, temp[1]})
			}
			if checkFresh(grid, temp[0]+1, temp[1]) {
				grid[temp[0]+1][temp[1]] = 3
				freshO--
				stack = append(stack, []int{temp[0] + 1, temp[1]})
			}
			if checkFresh(grid, temp[0], temp[1]-1) {
				grid[temp[0]][temp[1]-1] = 3
				freshO--
				stack = append(stack, []int{temp[0], temp[1] - 1})
			}
			if checkFresh(grid, temp[0], temp[1]+1) {
				grid[temp[0]][temp[1]+1] = 3
				freshO--
				stack = append(stack, []int{temp[0], temp[1] + 1})
			}
			stack = stack[1:]
		}
	}
	if freshO > 0 {
		return -1
	}
	return res
}

func checkFresh(grid [][]int, i, j int) bool {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && grid[i][j] == 1 {
		return true
	}
	return false
}
