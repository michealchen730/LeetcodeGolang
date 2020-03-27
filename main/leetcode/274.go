package main

import (
	"sort"
)

//leetcode题解解法
func hIndex_274(citations []int) int {
	sort.Ints(citations)
	res := 0
	for k, _ := range citations {
		if citations[len(citations)-k-1] >= k+1 {
			res = k + 1
		}
	}
	return res
}

//树状数组写法（太过于复杂）
func hIndex2(citations []int) int {
	arr := make([]int, len(citations)+2)
	res := 0
	for _, v := range citations {
		if v > len(citations) {
			v = len(citations)
		}
		updateRange(1, v, 1, arr)
	}
	for k, _ := range arr {
		if find(k, arr) >= k {
			res = k
		}
	}
	return res
}
func lowbit(x int) int {
	return x & (-x)
}
func update(i, k int, arr []int) {
	for i < len(arr) {
		arr[i] += k
		i += lowbit(i)
	}
}
func updateRange(i, j, k int, arr []int) {
	update(i, k, arr)
	update(j+1, -k, arr)
}
func find(i int, arr []int) int {
	res := 0
	for i > 0 {
		res += arr[i]
		i -= lowbit(i)
	}
	return res
}
