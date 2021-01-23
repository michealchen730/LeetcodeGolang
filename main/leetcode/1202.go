package main

import (
	"sort"
	"strings"
)

func smallestStringWithSwaps(s string, pairs [][]int) string {
	// 并查集
	// 按秩合并是必须的，否则会超时
	stu := make([]int, len(s))
	rank := make([]int, len(s))
	for k, _ := range stu {
		stu[k] = k
		rank[k] = 1
	}
	for _, v := range pairs {
		Bing(stu, rank, v[0], v[1])
	}
	mp := make(map[int][]byte)
	for k, _ := range s {
		if _, ok := mp[Cha(stu, k)]; ok {
			mp[Cha(stu, k)] = append(mp[Cha(stu, k)], s[k])
		} else {
			mp[Cha(stu, k)] = []byte{s[k]}
		}
	}
	for k, _ := range mp {
		sort.Sort(Bytes(mp[k]))
	}
	var res strings.Builder
	for k, _ := range s {
		res.WriteByte(mp[Cha(stu, k)][0])
		mp[Cha(stu, k)] = mp[Cha(stu, k)][1:]
	}
	return res.String()
}
