package main

func shiftGrid(grid [][]int, k int) [][]int {
	row := len(grid)
	col := len(grid[0])
	k = k % (row * col)
	t := k / col
	k = k % col
	grid = append(grid[row-t:], grid[0:row-t]...)
	for j := 0; j < k; j++ {
		t := grid[row-1][col-1]
		for n := row - 1; n > 0; n-- {
			grid[n] = append([]int{grid[n-1][col-1]}, grid[n][:col-1]...)
		}
		grid[0] = append([]int{t}, grid[0][:col-1]...)
	}
	return grid
}
