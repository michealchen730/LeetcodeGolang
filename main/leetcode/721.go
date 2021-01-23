package main

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	n := len(accounts)
	arr := make([]int, n)
	rank := make([]int, n)
	for i := range arr {
		arr[i] = i
		rank[i] = 1
	}
	mp := make(map[string]int)
	for k, v := range accounts {
		flag := true
		target := 0
		for idx, val := range v {
			if idx == 0 {
				continue
			}
			if m, ok := mp[val]; ok {
				Bing(arr, rank, m, k)
				flag = false
				target = Cha(arr, mp[val])
			}
		}
		if flag {
			for idx2, val := range v {
				if idx2 == 0 {
					continue
				}
				mp[val] = k
			}
		} else {
			for idex2, val := range v {
				if idex2 == 0 {
					continue
				}
				mp[val] = target
			}
		}
	}
	mp2 := make(map[int][]string)
	for k, v := range mp {
		mp2[Cha(arr, v)] = append(mp2[Cha(arr, v)], k)
	}

	var res [][]string
	for k, _ := range mp2 {
		sort.Strings(mp2[k])
		t := append([]string{accounts[k][0]}, mp2[k]...)
		res = append(res, t)
	}
	return res
}
