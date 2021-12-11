package main

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	if grid[row][col] != color {
		used := make([][]bool, len(grid))
		for k := range used {
			used[k] = make([]bool, len(grid[0]))
		}
		queue := [][]int{{row, col}}
		next := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for len(queue) != 0 {
			r, c := queue[0][0], queue[0][1]
			used[r][c] = true
			flag := false
			for _, v := range next {
				if 0 <= r+v[0] && r+v[0] < len(grid) && 0 <= c+v[1] && c+v[1] < len(grid[0]) {
					//r,c不是边界
					if !used[r+v[0]][c+v[1]] {
						if grid[r+v[0]][c+v[1]] != grid[r][c] {
							flag = true
						} else {
							queue = append(queue, []int{r + v[0], c + v[1]})
						}
					}
				} else {
					//r,c是边界
					flag = true
				}
			}
			if flag {
				grid[r][c] = color
			}
			queue = queue[1:]
		}
	}
	return grid
}
