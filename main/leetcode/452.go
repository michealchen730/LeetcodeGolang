package main

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Sort(Mat01(points))
	res := 0
	used := make([]bool, len(points))
	for k, v := range points {
		if !used[k] {
			res++
			used[k] = true
			right := v[1]
			for t := k + 1; t < len(points); t++ {
				if points[t][0] <= right {
					used[t] = true
					if points[t][1] < right {
						right = points[t][1]
					}
				} else {
					break
				}
			}
		}
	}
	return res
}
