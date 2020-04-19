package main

func uniquePathsIII(grid [][]int) int {
	left := 0
	sr, sc := 0, 0
	for k1, v1 := range grid {
		for k2, v2 := range v1 {
			if v2 == 0 {
				left++
			}
			if v2 == 1 {
				left++
				sr, sc = k1, k2
			}
		}
	}
	return dfs980(grid, sr, sc, left)
}
func dfs980(grid [][]int, i, j, left int) int {
	if grid[i][j] == 2 {
		if left == 0 {
			return 1
		} else {
			return 0
		}
	}
	if grid[i][j] == -1 {
		return 0
	}
	if grid[i][j] == 0 || grid[i][j] == 1 {
		temp := grid[i][j]
		grid[i][j] = -1
		p1, p2, p3, p4 := 0, 0, 0, 0
		if i+1 < len(grid) {
			p1 = dfs980(grid, i+1, j, left-1)
		}
		if i-1 >= 0 {
			p2 = dfs980(grid, i-1, j, left-1)
		}
		if j+1 < len(grid[0]) {
			p3 = dfs980(grid, i, j+1, left-1)
		}
		if j-1 >= 0 {
			p4 = dfs980(grid, i, j-1, left-1)
		}
		if temp != 1 {
			grid[i][j] = 0
		}
		return p1 + p2 + p3 + p4
	}
	return 0
}
