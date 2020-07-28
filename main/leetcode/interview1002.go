package main

import "sort"

func groupAnagramsI1002(strs []string) [][]string {
	mp := make(map[string][]string)
	for _, v := range strs {
		str := str2str(v)
		mp[str] = append(mp[str], v)
	}
	var res [][]string
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}

func str2str(s string) string {
	b := []byte(s)
	sort.Sort(Bytes(b))
	return string(b)
}
