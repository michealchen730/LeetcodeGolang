package main

func getMaximumGold(grid [][]int) int {
	res := []int{0}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			dfsMaxGold(grid, i, j, 0, res)
		}
	}
	return res[0]
}

func dfsMaxGold(grid [][]int, row, col, sum int, res []int) {
	if grid[row][col] != 0 {
		temp := grid[row][col]
		sum += temp
		grid[row][col] = 0
		f1, f2, f3, f4 := false, false, false, false
		if row > 0 && grid[row-1][col] != 0 {
			dfsMaxGold(grid, row-1, col, sum, res)
		} else {
			f1 = true
		}
		if row < len(grid)-1 && grid[row+1][col] != 0 {
			dfsMaxGold(grid, row+1, col, sum, res)
		} else {
			f2 = true
		}
		if col > 0 && grid[row][col-1] != 0 {
			dfsMaxGold(grid, row, col-1, sum, res)
		} else {
			f3 = true
		}
		if col < len(grid[0])-1 && grid[row][col+1] != 0 {
			dfsMaxGold(grid, row, col+1, sum, res)
		} else {
			f4 = true
		}
		if f1 && f2 && f3 && f4 && sum > res[0] {
			res[0] = sum
		}
		grid[row][col] = temp
	}
}
