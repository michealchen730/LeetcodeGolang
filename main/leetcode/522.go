package main

import (
	"sort"
)

func findLUSlength522(strs []string) int {
	sort.Sort(StringsByLength(strs))
	mp := make(map[string]int)
	var sts []string
	for i := 0; i < len(strs); i++ {
		mp[strs[i]]++
		if !(i > 0 && strs[i] == strs[i-1]) {
			sts = append(sts, strs[i])
		}
	}
	arr := make([]bool, len(sts))
	for i := len(sts) - 1; i >= 0; i-- {
		//重复出现的单词肯定不成立
		if mp[sts[i]] > 1 {
			arr[i] = true
		} else {
			flag := true
			for j := len(sts) - 1; j > i; j-- {
				//该单词是其他单词的子串，也不成立
				if longestCommonSubsequence(sts[i], sts[j]) == len(sts[i]) {
					arr[i] = true
					flag = false
					break
				}
			}
			if flag {
				return len(sts[i])
			}
		}
	}
	return -1
}
