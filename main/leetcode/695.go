package main

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	if len(grid) == 0 {
		return res
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				area := 0
				grid[i][j] = 2
				stack := [][]int{{i, j}}
				for len(stack) != 0 {
					temp := stack[0]
					stack = stack[1:]
					area++
					if checkIs(grid, temp[0]-1, temp[1]) {
						stack = append(stack, []int{temp[0] - 1, temp[1]})
					}
					if checkIs(grid, temp[0]+1, temp[1]) {
						stack = append(stack, []int{temp[0] + 1, temp[1]})
					}
					if checkIs(grid, temp[0], temp[1]-1) {
						stack = append(stack, []int{temp[0], temp[1] - 1})
					}
					if checkIs(grid, temp[0], temp[1]+1) {
						stack = append(stack, []int{temp[0], temp[1] + 1})
					}
				}
				if area > res {
					res = area
				}
			}
		}
	}
	return res
}

func checkIs(grid [][]int, i, j int) bool {
	if i >= 0 && i <= len(grid)-1 && j >= 0 && j <= len(grid[0])-1 && grid[i][j] == 1 {
		grid[i][j] = 2
		return true
	}
	return false
}
