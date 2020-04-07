package main

import (
	"strings"
)

type NodeW struct {
	word string
	fre  int
}

func topKFrequent2(words []string, k int) []string {
	mp := make(map[string]int)
	for _, v := range words {
		mp[v]++
	}
	var nodes []*NodeW
	for k, v := range mp {
		node := NodeW{
			word: k,
			fre:  v,
		}
		nodes = append(nodes, &node)
	}
	buildMinHeapK(nodes[:k])
	for i := k; i < len(nodes); i++ {
		if nodes[i].fre > nodes[0].fre || (nodes[i].fre == nodes[0].fre && strings.Compare(nodes[i].word, nodes[0].word) == -1) {
			nodes[0] = nodes[i]
			adjustMinHeapK(nodes, 0, k)
		}
	}
	var res []string
	for i := k - 1; i >= 0; i-- {
		res = append([]string{nodes[0].word}, res...)
		nodes[0] = nodes[i]
		adjustMinHeapK(nodes, 0, i)
	}
	return res
}
func buildMinHeapK(arr []*NodeW) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjustMinHeapK(arr, i, len(arr))
	}
}
func adjustMinHeapK(arr []*NodeW, i int, length int) {
	for k := 2*i + 1; k < length; k = 2*k + 1 {
		if k+1 < length && (arr[k].fre > arr[k+1].fre || (arr[k].fre == arr[k+1].fre && strings.Compare(arr[k+1].word, arr[k].word) != -1)) {
			k = k + 1
		}
		if arr[k].fre < arr[i].fre || (arr[k].fre == arr[i].fre && strings.Compare(arr[k].word, arr[i].word) != -1) {
			arr[k], arr[i] = arr[i], arr[k]
			i = k
		} else {
			break
		}
	}
}
