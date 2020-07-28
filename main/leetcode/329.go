package main

func longestIncreasingPath(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for k, _ := range dp {
		dp[k] = make([]int, len(matrix[0]))
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if dp[i][j] == 0 {
			dp[i][j] = 1
			if i-1 >= 0 && i-1 < len(dp) && matrix[i-1][j] < matrix[i][j] {
				dp[i][j] = max(dp[i][j], dfs(i-1, j)+1)
			}
			if i+1 >= 0 && i+1 < len(dp) && matrix[i+1][j] < matrix[i][j] {
				dp[i][j] = max(dp[i][j], dfs(i+1, j)+1)
			}
			if j-1 >= 0 && j-1 < len(dp[0]) && matrix[i][j-1] < matrix[i][j] {
				dp[i][j] = max(dp[i][j], dfs(i, j-1)+1)
			}
			if j+1 >= 0 && j+1 < len(dp[0]) && matrix[i][j+1] < matrix[i][j] {
				dp[i][j] = max(dp[i][j], dfs(i, j+1)+1)
			}
		}
		return dp[i][j]
	}
	res := 0
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			res = max(res, dfs(i, j))
		}
	}
	return res
}
