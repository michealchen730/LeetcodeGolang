package main

func numRookCaptures(board [][]byte) int {
	row, col, res := 0, 0, 0
FI:
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'R' {
				row, col = i, j
				break FI
			}
		}
	}
	for i := row - 1; i >= 0; i-- {
		if board[i][col] == 'p' {
			res++
			break
		} else if board[i][col] == 'B' {
			break
		}
	}
	for i := row + 1; i < len(board); i++ {
		if board[i][col] == 'p' {
			res++
			break
		} else if board[i][col] == 'B' {
			break
		}
	}
	for i := col + 1; i < len(board[0]); i++ {
		if board[row][i] == 'p' {
			res++
			break
		} else if board[row][i] == 'B' {
			break
		}
	}
	for i := col - 1; i >= 0; i-- {
		if board[row][i] == 'p' {
			res++
			break
		} else if board[row][i] == 'B' {
			break
		}
	}
	return res
}
