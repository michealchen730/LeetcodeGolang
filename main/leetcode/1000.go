package main

import (
	"math"
)

//以下均是参考题解

//解法一
func mergeStones(stones []int, K int) int {
	n := len(stones)
	if (n-1)%(K-1) != 0 {
		return -1
	}
	dp := make([][][]int, n+1)
	for k, _ := range dp {
		dp[k] = make([][]int, n+1)
		for i, _ := range dp[k] {
			dp[k][i] = make([]int, K+1)
			for m, _ := range dp[k][i] {
				dp[k][i][m] = math.MaxInt32
			}
		}
		dp[k][k][1] = 0
	}
	arr := make([]int, len(stones)+1)
	res := 0
	for i, v := range stones {
		res += v
		arr[i+1] = res
	}

	//定义区间长度
	for le := 2; le <= n; le++ {
		//枚举所有左边界
		for i := 1; i+le-1 <= n; i++ {
			//j表示右边界
			j := i + le - 1
			//枚举所有堆数
			for m := 2; m <= K; m++ {
				for p := i; p < j; p += K - 1 {
					dp[i][j][m] = min(dp[i][j][m], dp[i][p][1]+dp[p+1][j][m-1])
				}
			}
			dp[i][j][1] = dp[i][j][K] + arr[j] - arr[i-1]
		}
	}
	return dp[1][n][1]
}

//进一步压缩空间
func mergeStones2(stones []int, K int) int {
	if len(stones) == 1 {
		return 0
	}
	if len(stones) < K || (K != 2 && len(stones)%(K-1) != 1) {
		return -1
	}

	arr := make([]int, len(stones)+1)
	res := 0
	for i, v := range stones {
		res += v
		arr[i+1] = res
	}

	dp := make([][]int, len(stones)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(stones)+1)
	}
	//dp[i][j]表示从i到j的最大值

	for le := K; le <= len(stones); le++ {
		for i := 1; le+i-1 <= len(stones); i++ {
			j := le + i - 1
			dp[i][j] = math.MaxInt32

			for p := i; p < j; p += K - 1 {
				dp[i][j] = min(dp[i][j], dp[i][p]+dp[p+1][j])
			}
			if (j-i)%(K-1) == 0 {
				dp[i][j] += arr[j] - arr[i-1]
			}
		}
	}
	return dp[1][len(stones)]
}
