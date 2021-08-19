package main

import "math"

//可能起点和其他的点存在不互通的情况
//该算法在此种情况下会停止运算
func dijkstra(times [][]int, n int, k int) {
	//需要三个数组
	dist := make([]int, n)
	set := make([]bool, n)
	paths := make([]int, n)
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
			paths[idx] = -1
		}
	}
	paths[k-1] = k - 1
	for _, v := range mp[k-1] {
		dist[v[0]] = v[1]
		paths[v[0]] = k - 1
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
			break
		} else {
			set[nextTarget] = true
			for _, v := range mp[nextTarget] {
				if v[1]+dist[nextTarget] < dist[v[0]] {
					dist[v[0]] = v[1] + dist[nextTarget]
					paths[v[0]] = nextTarget
				}
			}
		}
		count++
	}
}
