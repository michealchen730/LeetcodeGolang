package main

func isValidSudoku2(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			used1 := make([]int, 9)
			used2 := make([]int, 9)
			if board[i][j] != '.' {
				if used1[board[i][j]-'1'] == 1 {
					return false
				} else {
					used1[board[i][j]-'1']++
				}
			}
			if board[j][i] != '.' {
				if used2[board[j][i]-'1'] == 1 {
					return false
				} else {
					used2[board[j][i]-'1']++
				}
			}
		}
	}
	for i := 1; i <= 7; i += 3 {
		for j := 1; j <= 7; j += 3 {
			used3 := make([]int, 9)
			for k := -1; k <= 1; k++ {
				for m := -1; m <= 1; m++ {
					if board[i+k][j+m] != '.' {
						if used3[board[i+k][j+m]-'1'] == 1 {
							return false
						} else {
							used3[board[i+k][j+m]-'1']++
						}
					}
				}
			}
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	matrix := make([][]int, 27)
	for i, _ := range matrix {
		matrix[i] = make([]int, 9)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] != '.' {
				temp := board[i][j] - '1'
				if matrix[i][temp] == 1 || matrix[9+j][temp] == 1 || matrix[18+i/3*3+j/3][temp] == 1 {
					return false
				} else {
					matrix[i][temp], matrix[9+j][temp], matrix[18+i/3*3+j/3][temp] = 1, 1, 1
				}
			}
		}
	}
	return true
}
