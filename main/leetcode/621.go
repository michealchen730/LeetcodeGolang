package main

import (
	"sort"
)

type Ints2 []int

func (b Ints2) Len() int           { return len(b) }
func (b Ints2) Less(i, j int) bool { return b[i] > b[j] }
func (b Ints2) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func leastInterval(tasks []byte, n int) int {
	arr := make([]int, 26)
	for _, v := range tasks {
		arr[v-'A']++
	}
	res := 0
	sort.Sort(Ints2(arr))
	for !(arr[0] == 1 && arr[min(n+1, 25)] <= 0) {
		res += n + 1
		for i := 0; i < min(n+1, 25); i++ {
			arr[i]--
		}
		sort.Sort(Ints2(arr))
	}
	temp := 0
	for _, v := range arr {
		if v == 1 {
			temp++
		}
	}
	return res + temp
}
