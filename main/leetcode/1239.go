package main

func maxLength(arr []string) int {
	//单词本身可能就含有重复字母
	var cand []string
	for _, v := range arr {
		mp := make(map[rune]bool)
		flag := true
		for _, r := range v {
			if mp[r] {
				flag = false
				break
			}
		}
		if flag {
			cand = append(cand, v)
		}
	}
	if len(cand) == 0 {
		return 0
	}
	//建立每个单词可去的地方
	paths := make([][]int, len(cand))
	for idx, _ := range paths {
		paths[idx] = []int{}
	}
	for k, word1 := range cand {
		mp := make(map[rune]bool)
		for _, r := range word1 {
			mp[r] = true
		}
		for idx := k + 1; idx < len(cand); idx++ {
			flag := true
			for _, r := range cand[idx] {
				if mp[r] {
					break
				}
			}
			if flag {
				paths[k] = append(paths[k], idx)
			}
		}
	}

	getNewPath := func(a, b []int) []int {
		mp := make(map[int]bool)
		for _, v := range b {
			mp[v] = true
		}
		var res []int
		for _, v := range a {
			if mp[v] {
				res = append(res, v)
			}
		}
		return res
	}

	res := 0
	//dfs
	var dfs func(i, t int, path []int)

	dfs = func(idx, tlh int, path []int) {
		tlh += len(cand[idx])
		if tlh > res {
			res = tlh
		}
		//遍历当前能走的每一条路径
		for _, v := range path {
			newPath := getNewPath(path, paths[v])
			if len(newPath) == 0 {
				return
			}
			dfs(v, tlh, newPath)
		}
	}

	for i := 0; i < len(cand); i++ {
		dfs(i, 0, paths[i])
	}
	return res
}
