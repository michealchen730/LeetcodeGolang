package main

func tictactoe(moves [][]int) string {
	if len(moves) < 5 {
		return "Pending"
	}
	var grid [3][3]int
	for k, v := range moves {
		if k%2 == 0 {
			grid[v[0]][v[1]] = 1
		} else {
			grid[v[0]][v[1]] = 2
		}
		if k >= 4 {
			if checkGrid(grid) {
				if k%2 == 0 {
					return "A"
				} else {
					return "B"
				}
			}
		}
	}
	if len(moves) < 9 {
		return "Pending"
	} else {
		return "Draw"
	}
}

func checkGrid(grid [3][3]int) bool {
	for i := 0; i < 3; i++ {
		if grid[0][i] != 0 && grid[0][i] == grid[1][i] && grid[1][i] == grid[2][i] {
			return true
		}
		if grid[i][0] != 0 && grid[i][0] == grid[i][1] && grid[i][1] == grid[i][2] {
			return true
		}
	}
	if grid[1][1] != 0 {
		if grid[0][0] == grid[1][1] && grid[1][1] == grid[2][2] {
			return true
		}
		if grid[2][0] == grid[1][1] && grid[1][1] == grid[0][2] {
			return true
		}
	}
	return false
}
