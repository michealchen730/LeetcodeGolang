package main

func islandPerimeter(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				res += 4
				if i >= 1 && grid[i-1][j] == 1 {
					res -= 1
				}
				if i+1 < len(grid) && grid[i+1][j] == 1 {
					res -= 1
				}
				if j >= 1 && grid[i][j-1] == 1 {
					res -= 1
				}
				if j+1 < len(grid[0]) && grid[i][j+1] == 1 {
					res -= 1
				}
			}

		}
	}
	return res
}
