package main

import (
	"math"
	"sort"
)

func minDistance1478(houses []int, k int) int {
	sort.Ints(houses)
	n := len(houses)
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, k+1)
	}
	cost := make([][]int, n)
	for i, _ := range cost {
		cost[i] = make([]int, n)
	}

	//cost[i][j]表示第i栋房子和第j栋房子之间架设一个邮桶后，中间每个房子到这个邮筒的距离和
	//先看这个
	//[4,8,10]三栋房子，邮筒是否应该架设7的位置呢，答案是否定的，应该架设在8的位置上
	//因此，对于这个放置问题，可以总结为：
	//当只有单数个房子时，邮筒放在最中间那个房子的位置
	//当只有偶数个房子时，邮筒放在中间两个房子的任一一个即可
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			for k := i; k <= j; k++ {
				cost[i][j] += abs(houses[k] - houses[i+(j-i+1)/2])
			}
		}
	}

	for i := 0; i < n; i++ {
		dp[i][0] = math.MaxInt32
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j] = math.MaxInt32
			//这里是一次枚举，枚举所有可能的i
			for m := 0; m <= i; m++ {
				t := 0
				if m > 0 {
					t = dp[m-1][j-1]
				}
				dp[i][j] = min(dp[i][j], t+cost[m][i])
			}
		}
	}
	return dp[n-1][k]
}
