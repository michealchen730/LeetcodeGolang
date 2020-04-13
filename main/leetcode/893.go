package main

import (
	"sort"
)

func numSpecialEquivGroups(A []string) int {
	res := 0
	mp := make(map[string]int)
	for _, v := range A {
		var b1 []byte
		var b2 []byte
		for k, _ := range v {
			if k%2 == 0 {
				b1 = append(b1, v[k])
			} else {
				b2 = append(b2, v[k])
			}
		}
		sort.Sort(Bytes(b1))
		sort.Sort(Bytes(b2))
		s := string(b1) + string(b2)
		if _, ok := mp[s]; !ok {
			res++
			mp[s] = 1
		}
	}
	return res
}
