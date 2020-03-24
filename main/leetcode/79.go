package main

func exist(board [][]byte, word string) bool {
	if len(board)*len(board[0]) < len(word) {
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if exist2(board, word, i, j) {
				return true
			}
		}
	}
	return false
}

func exist2(board [][]byte, word string, i, j int) bool {
	if word == "" {
		return true
	}
	if !(i >= 0 && i < len(board) && j >= 0 && j < len(board[0])) {
		return false
	}
	if board[i][j] == word[0] {
		board[i][j] = '1'
		if exist2(board, word[1:], i+1, j) || exist2(board, word[1:], i-1, j) || exist2(board, word[1:], i, j+1) || exist2(board, word[1:], i, j-1) {
			return true
		}
		board[i][j] = word[0]
	}
	return false
}
