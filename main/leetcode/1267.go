package main

func countServers(grid [][]int) int {
	mp1 := make(map[int]int)
	mp2 := make(map[int]int)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				mp1[i]++
				mp2[j]++
			}
		}
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 && (mp1[i] > 1 || mp2[j] > 1) {
				res++
			}
		}
	}
	return res
}
