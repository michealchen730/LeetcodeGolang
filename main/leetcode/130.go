package main

func solve(board [][]byte) {
	if len(board) <= 2 || len(board[0]) <= 2 {
		return
	}
	var stack [][]int
	for i := 0; i < len(board); i++ {
		if board[i][0] == 'O' {
			board[i][0] = 'U'
			stack = append(stack, []int{i, 0})
		}
		if board[i][len(board[0])-1] == 'O' {
			board[i][len(board[0])-1] = 'U'
			stack = append(stack, []int{i, len(board[0]) - 1})
		}
	}
	for i := 0; i < len(board[0]); i++ {
		if board[0][i] == 'O' {
			board[0][i] = 'U'
			stack = append(stack, []int{0, i})
		}
		if board[len(board)-1][i] == 'O' {
			board[len(board)-1][i] = 'U'
			stack = append(stack, []int{len(board) - 1, i})
		}
	}
	for len(stack) != 0 {
		temp := stack[0]
		if checkO(board, temp[0]-1, temp[1]) {
			board[temp[0]-1][temp[1]] = 'U'
			stack = append(stack, []int{temp[0] - 1, temp[1]})
		}
		if checkO(board, temp[0]+1, temp[1]) {
			board[temp[0]+1][temp[1]] = 'U'
			stack = append(stack, []int{temp[0] + 1, temp[1]})
		}
		if checkO(board, temp[0], temp[1]+1) {
			board[temp[0]][temp[1]+1] = 'U'
			stack = append(stack, []int{temp[0], temp[1] + 1})
		}
		if checkO(board, temp[0], temp[1]-1) {
			board[temp[0]][temp[1]-1] = 'U'
			stack = append(stack, []int{temp[0], temp[1] - 1})
		}
		stack = stack[1:]
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'U' {
				board[i][j] = 'O'
			}
		}
	}
}

func checkO(board [][]byte, i int, j int) bool {
	if i >= 1 && i <= len(board)-2 && j >= 1 && j <= len(board[0])-2 && board[i][j] == 'O' {
		return true
	}
	return false
}
