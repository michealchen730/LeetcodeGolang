package main

import (
	"sort"
)

type S953 struct {
	strs []string
	mp   map[byte]int
}

func (b S953) Len() int { return len(b.strs) }
func (b S953) Less(i, j int) bool {
	k := 0
	for k < len(b.strs[i]) && k < len(b.strs[j]) {
		if b.strs[i][k] == b.strs[j][k] {
			k++
			continue
		}
		if b.mp[b.strs[i][k]] < b.mp[b.strs[j][k]] {
			return true
		}
		if b.mp[b.strs[i][k]] > b.mp[b.strs[j][k]] {
			return false
		}
	}
	return len(b.strs[i]) < len(b.strs[j])
}
func (b S953) Swap(i, j int) { b.strs[i], b.strs[j] = b.strs[j], b.strs[i] }

func isAlienSorted(words []string, order string) bool {
	cpy := make([]string, len(words))
	copy(cpy, words)
	mp := make(map[byte]int)
	for k, _ := range order {
		mp[order[k]] = k
	}
	sort.Sort(S953{strs: cpy, mp: mp})
	for i := 0; i < len(cpy); i++ {
		if cpy[i] != words[i] {
			return false
		}
	}
	return true
}
