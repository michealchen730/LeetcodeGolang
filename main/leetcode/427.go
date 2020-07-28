package main

type Node4 struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node4
	TopRight    *Node4
	BottomLeft  *Node4
	BottomRight *Node4
}

func construct(grid [][]int) *Node4 {
	flag := grid[0][0]
FI:
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != grid[0][0] {
				flag = 2
				break FI
			}
		}
	}
	if flag != 2 {
		b := true
		if flag == 0 {
			b = false
		}
		return &Node4{
			Val:    b,
			IsLeaf: true,
		}
	} else {
		row := len(grid)
		col := len(grid[0])
		var tleft [][]int
		var tright [][]int
		var bleft [][]int
		var bright [][]int
		for i := 0; i < row/2; i++ {
			tleft = append(tleft, grid[i][0:col/2])
			tright = append(tright, grid[i][col/2:])
		}
		for i := row / 2; i < row; i++ {
			bleft = append(bleft, grid[i][0:col/2])
			bright = append(bright, grid[i][col/2:])
		}

		return &Node4{
			Val:         true,
			IsLeaf:      false,
			TopLeft:     construct(tleft),
			TopRight:    construct(tright),
			BottomLeft:  construct(bleft),
			BottomRight: construct(bright),
		}
	}
}
