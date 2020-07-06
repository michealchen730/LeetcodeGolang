package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	r, c := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp := make([][]int, r)
	for k, _ := range dp {
		dp[k] = make([]int, c)
	}
	dp[0][0] = 1
	for i := 1; i < r; i++ {
		if obstacleGrid[i][0] == 1 {
			dp[i][0] = 0
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}
	for j := 1; j < c; j++ {
		if obstacleGrid[0][j] == 1 {
			dp[0][j] = 0
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[r-1][c-1]
}
