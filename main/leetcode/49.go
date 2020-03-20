package main

import "sort"

type Bytes []byte

func (b Bytes) Len() int           { return len(b) }
func (b Bytes) Less(i, j int) bool { return b[i] < b[j] }
func (b Bytes) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)
	for _, v := range strs {
		bytes := []byte(v)
		sort.Sort(Bytes(bytes))
		s := string(bytes)
		mp[s] = append(mp[s], v)
	}
	var res [][]string
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}
