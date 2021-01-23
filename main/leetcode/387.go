package main

import (
	"math"
)

func firstUniqChar387(s string) int {
	res := math.MaxInt32
	//记录该字母是否重复
	repeated := make([]int, 26)
	mp := make(map[int]int)
	for k, _ := range s {
		repeated[int(s[k]-'a')]++
		//哈希表用来记录每个字母出现的最早位置
		if _, ok := mp[int(s[k]-'a')]; !ok {
			mp[int(s[k]-'a')] = k
		}
	}
	flag := false
	for k, v := range repeated {
		if v == 1 {
			flag = true
			res = min(res, mp[k])
		}
	}
	if flag {
		return res
	}
	return -1
}
