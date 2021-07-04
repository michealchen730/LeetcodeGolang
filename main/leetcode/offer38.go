package main

import "sort"

func permutation38(s string) []string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	var path []byte
	used := make([]bool, len(b))
	var res []string
	var dfs func()
	dfs = func() {
		if len(path) == len(b) {
			res = append(res, string(path))
			return
		}

		for k, v := range used {
			if !v {
				//有效位的剪枝
				if k > 0 && b[k] == b[k-1] && used[k-1] == false {
					continue
				}
				path = append(path, b[k])
				used[k] = true
				dfs()
				used[k] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs()
	return res

}
