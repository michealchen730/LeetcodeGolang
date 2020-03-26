package main

func gameOfLife(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			temp := countCellsAlive(board, i, j)
			if board[i][j] == 0 {
				if temp == 3 {
					board[i][j] = -1
				}
			} else {
				if temp < 2 || temp > 3 {
					board[i][j] = 2
				}
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 2 {
				board[i][j] = 0
			}
			if board[i][j] == -1 {
				board[i][j] = 1
			}
		}
	}
}
func countCellsAlive(board [][]int, i, j int) int {
	res := 0
	for m := -1; m <= 1; m++ {
		for n := -1; n <= 1; n++ {
			if 0 <= (i+m) && (i+m) < len(board) && (j+n) >= 0 && (j+n) < len(board[0]) && board[i+m][j+n] > 0 {
				res++
			}
		}
	}
	if board[i][j] == 1 {
		res--
	}
	return res
}
