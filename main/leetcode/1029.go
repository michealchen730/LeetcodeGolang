package main

import (
	"sort"
)

type intS [][]int

func (s intS) Len() int           { return len(s) }
func (s intS) Less(i, j int) bool { return s[i][2] > s[j][2] }
func (s intS) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func twoCitySchedCost(costs [][]int) int {
	if len(costs) <= 1 {
		return 0
	}
	for i := 0; i < len(costs); i++ {
		costs[i] = append(costs[i], abs(costs[i][0]-costs[i][1]))
	}
	sort.Sort(intS(costs))
	toA, toB, sum := 0, 0, 0
	for i := 0; i < len(costs); i++ {
		if toA == len(costs)/2 {
			sum += costs[i][1]
			continue
		}
		if toB == len(costs)/2 {
			sum += costs[i][0]
			continue
		}
		if costs[i][0] > costs[i][1] {
			sum += costs[i][1]
			toB++
		} else {
			sum += costs[i][0]
			toA++
		}
	}
	return sum
}

func abs(i int) int {
	if i < 0 {
		i = -i
	}
	return i
}
