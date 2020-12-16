package main

import (
	"sort"
)

//从起始点的动态规划
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	mx := 1

	sort.Sort(Mat01(intervals))

	dp := make([]int, len(intervals))
	dp[0] = 1
	for i := 1; i < len(intervals); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if intervals[i][0] >= intervals[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		mx = max(mx, dp[i])
	}

	return len(intervals) - mx

}

//从终点的动态规划
//这里就类似于更直观的背包问题了
//intervals的每一个元素都相当于一件物品，可以选择放或者不放
//dp[i]代表到第i个区间的最大区间数
func eraseOverlapIntervals2(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	mx := 1

	sort.Sort(Mat01(intervals))

	dp := make([]int, len(intervals))
	dp[0] = 1
	for i := 1; i < len(intervals); i++ {
		//如果不要当前区间，那么dp[i-1]就是最大区间
		dp[i] = dp[i-1]
		for j := i - 1; j >= 0; j-- {
			if intervals[i][0] >= intervals[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
				break
			}
		}
		mx = max(mx, dp[i])
	}

	return len(intervals) - mx

}
