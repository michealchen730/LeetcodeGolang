package main

func surfaceArea(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 0 {
				res += (grid[i][j]-1)*4 + 6
				if i != 0 && grid[i-1][j] != 0 {
					res -= min(grid[i-1][j], grid[i][j]) * 2
				}
				if j != 0 && grid[i][j-1] != 0 {
					res -= min(grid[i][j], grid[i][j-1]) * 2
				}
			}
		}
	}
	return res
}
