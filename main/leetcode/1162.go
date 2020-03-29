package main

func maxDistance(grid [][]int) int {
	var islands [][]int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				islands = append(islands, []int{i, j})
			}
		}
	}
	if len(islands) == 0 || len(islands) == len(grid)*len(grid[0]) {
		return -1
	}
	dis := -1
	for len(islands) != 0 {
		temp := islands
		islands = [][]int{}
		for _, v := range temp {
			if checkIsland(grid, v[0]+1, v[1]) == 0 {
				grid[v[0]+1][v[1]] = 1
				islands = append(islands, []int{v[0] + 1, v[1]})
			}
			if checkIsland(grid, v[0]-1, v[1]) == 0 {
				grid[v[0]-1][v[1]] = 1
				islands = append(islands, []int{v[0] - 1, v[1]})
			}
			if checkIsland(grid, v[0], v[1]+1) == 0 {
				grid[v[0]][v[1]+1] = 1
				islands = append(islands, []int{v[0], v[1] + 1})
			}
			if checkIsland(grid, v[0], v[1]-1) == 0 {
				grid[v[0]][v[1]-1] = 1
				islands = append(islands, []int{v[0], v[1] - 1})
			}
		}
		dis++
	}
	return dis
}

func checkIsland(grid [][]int, i, j int) int {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) {
		return grid[i][j]
	} else {
		return -1
	}
}
