package main

func snakesAndLadders(board [][]int) int {
	N := len(board)
	used := make(map[int]bool)
	transform := func(index int) (int, int) {
		r := (index - 1) / N
		c := (index - 1) % N
		if r%2 != 0 {
			c = N - 1 - c
		}
		return N - 1 - r, c
	}
	queue := []int{1}
	step := 0
	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			e := queue[i]
			if used[e] {
				continue
			}
			used[e] = true
			if e == N*N {
				return step
			}
			for j := e + 1; j <= min(e+6, N*N); j++ {
				r, c := transform(j)
				if board[r][c] != -1 {
					queue = append(queue, board[r][c])
				} else {
					queue = append(queue, j)
				}
			}
		}
		queue = queue[l:]
		step += 1
	}
	return -1
}
