package main

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] <= envelopes[j][1]
		} else {
			return false
		}
	})
	res := 0
	dp := make([]int, len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if envelopes[i][0] > envelopes[j][0] && envelopes[i][1] > envelopes[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

//基于二分查找的动态规划
//这题的正确性证明很有意思
//因为每一次append操作都是正确的，所以最后f的长度也是正确的
//但对于f[searchInts]=envelopes[i][1]，则不保证正确性
//比如[[1,12],[2,18],[6,19],[7,13]]
func maxEnvelopes2(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1]
	})
	f := []int{}
	for i := 0; i < len(envelopes); i++ {

		if searchInts := sort.SearchInts(f, envelopes[i][1]); searchInts < len(f) {
			f[searchInts] = envelopes[i][1]
		} else {
			f = append(f, envelopes[i][1])
		}
	}
	return len(f)
}
