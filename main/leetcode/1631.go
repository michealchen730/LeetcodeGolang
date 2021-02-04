package main

import (
	"sort"
)

func minimumEffortPath(heights [][]int) int {
	//m*n个点
	m, n := len(heights), len(heights[0])
	l := m * n
	root := make([]int, l+1)
	rank := make([]int, l+1)
	for k, _ := range root {
		root[k] = k
		rank[k] = 1
	}
	//将边抽象出来
	var edges [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i+1 < m {
				edges = append(edges, []int{abs(heights[i][j] - heights[i+1][j]), i*n + j + 1, (i+1)*n + j + 1})
			}
			if j+1 < n {
				edges = append(edges, []int{abs(heights[i][j] - heights[i][j+1]), i*n + j + 1, i*n + j + 2})
			}
		}
	}
	sort.Sort(Edges(edges))
	for _, v := range edges {
		Bing(root, rank, v[1], v[2])
		if Cha(root, 1) == Cha(root, m*n) {
			return v[0]
		}
	}
	return 0
}

type Edges [][]int

func (s Edges) Len() int           { return len(s) }
func (s Edges) Less(i, j int) bool { return s[i][0] < s[j][0] }
func (s Edges) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
