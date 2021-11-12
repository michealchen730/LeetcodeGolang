package main

import (
	"container/heap"
)

type grids [][]int

func (g grids) Len() int            { return len(g) }
func (g grids) Swap(i, j int)       { g[i], g[j] = g[j], g[i] }
func (g grids) Less(i, j int) bool  { return g[i][2] < g[j][2] }
func (g *grids) Push(v interface{}) { *g = append(*g, v.([]int)) }
func (g *grids) Pop() interface{}   { a := *g; v := a[len(a)-1]; *g = a[:len(a)-1]; return v }

func trapRainWater(heightMap [][]int) int {
	m, n := len(heightMap), len(heightMap[0])
	if m < 3 || n < 3 {
		return 0
	}
	h := &grids{}
	used := make([][]int, m)
	for k := range used {
		used[k] = make([]int, n)
	}
	for row, arr := range heightMap {
		for col, v := range arr {
			if row == 0 || row == m-1 || col == 0 || col == n-1 {
				heap.Push(h, []int{row, col, v})
				used[row][col] = 1
			}
		}
	}
	next := []int{-1, 0, 1, 0, -1}
	ans := 0
	for h.Len() > 0 {
		c := heap.Pop(h).([]int)
		for i := 0; i < 4; i++ {
			cx, cy := c[0]+next[i], c[1]+next[i+1]
			if cx >= 0 && cx < m && cy >= 0 && cy < n && used[cx][cy] == 0 {
				if heightMap[cx][cy] < c[2] {
					ans += c[2] - heightMap[cx][cy]
					heap.Push(h, []int{cx, cy, c[2]})
				} else {
					heap.Push(h, []int{cx, cy, heightMap[cx][cy]})
				}
				used[cx][cy] = 1
			}
		}
	}
	return ans
}
