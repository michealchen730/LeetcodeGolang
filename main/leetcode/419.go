package main

func countBattleships(board [][]byte) int {
	res := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'X' {
				if (i >= 1 && board[i-1][j] != 'X') || (j >= 1 && board[i][j-1] == 'X') {
					continue
				}
				res++
			}
		}
	}
	return res
}
