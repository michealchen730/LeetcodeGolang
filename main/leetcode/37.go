package main

//这也是我不会的题目
//原装算法如下
func solveSudoku(board [][]byte) {
	//rows，cols和boxes就是代表的三个约束
	rows := [9][10]int{}
	cols := [9][10]int{}
	boxes := [9][10]int{}

	sudokuSolved := false

	var backtrack func(row, col int)
	var placeNextNumber func(row, col int)

	//能否放置数字
	couldPlace := func(val, row, col int) bool {
		//如果我们将数独的区域划分成九块，如何定义哪一行哪一列在哪一块中
		idx := (row/3)*3 + col/3
		return rows[row][val]+cols[col][val]+boxes[idx][val] == 0
	}

	//放置数字
	placeNumber := func(val, row, col int) {
		idx := (row/3)*3 + col/3
		rows[row][val]++
		cols[col][val]++
		boxes[idx][val]++
		board[row][col] = byte(val + '0')
	}

	//移除数字
	removeNumber := func(val, row, col int) {
		idx := (row/3)*3 + col/3
		rows[row][val]--
		cols[col][val]--
		boxes[idx][val]--
		board[row][col] = '.'
	}

	//放置下一个数字
	//这里应该是回溯的主要部分
	placeNextNumber = func(row, col int) {
		if col == 8 && row == 8 {
			sudokuSolved = true
		} else {
			if col == 8 {
				backtrack(row+1, 0)
			} else {
				backtrack(row, col+1)
			}
		}
	}

	backtrack = func(row, col int) {
		if board[row][col] == '.' {
			//如果某行某列所有val全部试了，那么不会继续调用placeNextNumber，算法的递归结束
			for val := 1; val < 10; val++ {
				if couldPlace(val, row, col) {
					placeNumber(val, row, col)
					//placeNextNumber会试出所有数字
					placeNextNumber(row, col)
					if !sudokuSolved {
						removeNumber(val, row, col)
					}
				}
			}
		} else {
			placeNextNumber(row, col)
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				val := board[i][j] - '0'
				placeNumber(int(val), i, j)
			}
		}
	}
	backtrack(0, 0)
}
