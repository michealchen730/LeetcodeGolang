package main

func cherryPickup(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([][][]int, n)
	for k, _ := range dp {
		dp[k] = make([][]int, m)
		for i, _ := range dp[k] {
			dp[k][i] = make([]int, m)
			for j, _ := range dp[k][i] {
				dp[k][i][j] = -1
			}
		}
	}
	dp[0][0][len(grid[0])-1] = grid[0][0] + grid[0][len(grid[0])-1]
	for i := 0; i < n-1; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < m; k++ {
				if dp[i][j][k] != -1 {
					for s := -1; s <= 1; s++ {
						for u := -1; u <= 1; u++ {
							if j+s >= 0 && j+s < m && k+u >= 0 && k+u < m {
								t := grid[i+1][j+s]
								if j+s != k+u {
									t += grid[i+1][k+u]
								}
								dp[i+1][j+s][k+u] = max(dp[i+1][j+s][k+u], t+dp[i][j][k])
							}
						}
					}
				}
			}
		}
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			res = max(res, dp[n-1][i][j])
		}
	}
	return res
}
