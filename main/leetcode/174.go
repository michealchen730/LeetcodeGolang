package main

func calculateMinimumHP(dungeon [][]int) int {
	//dp[i][j]表示走到dp[i][j]需要的最少生命值
	dp := make([][]int, len(dungeon))
	for k, _ := range dp {
		dp[k] = make([]int, len(dungeon[0]))
	}
	//边界问题处理
	dp[len(dp)-1][len(dp[0])-1] = max(0-dungeon[len(dungeon)-1][len(dungeon[0])-1]+1, 1)
	for i := len(dp) - 2; i >= 0; i-- {
		dp[i][len(dp[0])-1] = max(dp[i+1][len(dp[0])-1]-dungeon[i][len(dungeon[0])-1], 1)
	}
	for i := len(dp[0]) - 2; i >= 0; i-- {
		dp[len(dp)-1][i] = max(dp[len(dp)-1][i+1]-dungeon[len(dungeon)-1][i], 1)
	}
	for i := len(dp) - 2; i >= 0; i-- {
		for j := len(dp[0]) - 2; j >= 0; j-- {
			dp[i][j] = max(min(dp[i+1][j], dp[i][j+1])-dungeon[i][j], 1)
		}
	}
	return dp[0][0]
}
