package main

import (
	"math"
)

func networkDelayTime(times [][]int, n int, k int) int {
	//需要三个数组
	dist := make([]int, n)
	set := make([]bool, n)
	//需要初始化图
	mp := make(map[int][][]int)
	for _, v := range times {
		if _, ok := mp[v[0]-1]; !ok {
			mp[v[0]-1] = [][]int{{v[1] - 1, v[2]}}
		} else {
			mp[v[0]-1] = append(mp[v[0]-1], []int{v[1] - 1, v[2]})
		}
	}
	//初始化三个数组
	for idx, _ := range dist {
		if idx != k-1 {
			dist[idx] = math.MaxInt32
		}
	}
	for _, v := range mp[k-1] {
		dist[v[0]] = v[1]
	}
	set[k-1] = true

	count := 1
	for count < n {
		nextTarget := -1
		for idx, v := range set {
			if !v && dist[idx] != math.MaxInt32 {
				if nextTarget == -1 {
					nextTarget = idx
				} else {
					if dist[idx] < dist[nextTarget] {
						nextTarget = idx
					}
				}
			}
		}
		if nextTarget == -1 {
			return -1
		} else {
			set[nextTarget] = true
			for _, v := range mp[nextTarget] {
				if v[1]+dist[nextTarget] < dist[v[0]] {
					dist[v[0]] = v[1] + dist[nextTarget]
				}
			}
		}
		count++
	}
	res := 0
	for _, v := range dist {
		res = max(res, v)
	}
	return res
}
