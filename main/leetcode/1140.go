package main

func stoneGameII(piles []int) int {
	//dp[i][j]表示在第i堆，M=j的情况下能拿到的最大值（并不是这一轮的）
	dp := make([][]int, len(piles))
	for k, _ := range dp {
		dp[k] = make([]int, len(piles)+1)
	}
	//零和博弈的题多半是从后往前dp
	sum := 0
	for i := len(dp) - 1; i >= 0; i-- {
		sum += piles[i]
		for M := 1; M <= len(piles); M++ {
			if i+2*M >= len(piles) { //边界值的情况，举个例子，剩余两个石堆，如果我们能拿两个，那么我们肯定全拿
				dp[i][M] = sum
			} else {
				for x := 1; i+x <= len(piles) && x <= 2*M; x++ {
					dp[i][M] = max(dp[i][M], sum-dp[i+x][max(M, x)])
				}
			}
		}
	}
	return dp[0][1]
}
