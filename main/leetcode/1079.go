package main

import "sort"

type Bytes1079 []byte

func (b Bytes1079) Len() int           { return len(b) }
func (b Bytes1079) Less(i, j int) bool { return b[i] < b[j] }
func (b Bytes1079) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func numTilePossibilities(tiles string) int {
	bytes := []byte(tiles)
	sort.Sort(Bytes1079(bytes))
	used := make([]int, len(bytes))
	arr := []int{0}
	dfs1079(0, bytes, used, arr)
	//减一的原因是空集也被纳入集合中去了
	return arr[0] - 1
}

func dfs1079(i int, bytes []byte, used, arr []int) {
	//每取一次就加一
	arr[0]++
	//终止
	if i == len(bytes) {
		return
	}
	for j := 0; j < len(used); j++ {
		if used[j] == 0 && ((j == 0) || (j > 0 && (bytes[j] != bytes[j-1] || (bytes[j] == bytes[j-1] && used[j-1] == 1)))) {
			used[j] = 1
			dfs1079(i+1, bytes, used, arr)
			used[j] = 0
		}
	}
}
