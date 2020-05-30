package main

import (
	"strconv"
	"strings"
)

func largestNumber(cost []int, target int) string {
	mp := make(map[int]int)
	for k, v := range cost {
		mp[v] = max(mp[v], k+1)
	}
	dp := make([]string, target+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = "nil"
	}
	for i := 1; i < len(dp); i++ {
		if mp[i] != 0 {
			dp[i] = strconv.Itoa(mp[i])
		}
		for k, v := range mp {
			if i-k > 0 && dp[i-k] != "nil" {
				var res strings.Builder
				res.WriteString(dp[i-k])
				res.WriteByte(byte(v + 48))
				dp[i] = maxStrings(dp[i], res.String())
			}
		}
	}
	if dp[target] == "nil" {
		return "0"
	}
	return dp[target]
}

func maxStrings(s1, s2 string) string {
	if s1 == "nil" {
		return s2
	}
	if len(s1) > len(s2) {
		return s1
	} else if len(s1) < len(s2) {
		return s2
	} else {
		for i := 0; i < len(s1); i++ {
			if s1[i] > s2[i] {
				return s1
			} else if s1[i] < s2[i] {
				return s2
			}
		}
	}
	return " "
}
