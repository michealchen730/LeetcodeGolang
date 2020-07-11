package main

import "sort"

func permutationI0808(S string) []string {
	var res []string
	var path []byte
	bytes := []byte(S)
	sort.Sort(Bytes1079(bytes))
	used := make([]int, len(S))
	dfsI0808(bytes, &path, used, &res)
	return res
}

func dfsI0808(arr []byte, path *[]byte, used []int, res *[]string) {
	if len(*path) == len(arr) {
		s := string(*path)
		*res = append(*res, s)
		return
	}
	for i := 0; i < len(arr); i++ {
		if used[i] == 0 {
			if i > 0 && arr[i] == arr[i-1] && used[i-1] == 0 {
				continue
			}
			*path = append(*path, arr[i])
			used[i] = 1
			dfsI0808(arr, path, used, res)
			used[i] = 0
			*path = (*path)[:len(*path)-1]
		}
	}
}
