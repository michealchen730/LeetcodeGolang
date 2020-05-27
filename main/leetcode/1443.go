package main

func minTime(n int, edges [][]int, hasApple []bool) int {
	//重建图
	mp := make(map[int][]int)
	for _, v := range edges {
		mp[v[0]] = append(mp[v[0]], v[1])
		mp[v[1]] = append(mp[v[1]], v[0])
	}
	length := 0
	showed := make(map[int]int)
	for i := 0; i < len(hasApple); i++ {
		if hasApple[i] == true {
			used := make(map[int]int)
			used[i] = 1
			path := []int{i}
			dfs1443(mp, &path, i, used)
			if len(path) != 0 {
				for k, v := range path {
					if showed[v] == 0 {
						showed[v] = 1
						if k != 0 {
							length += 2
						}
					} else {
						length += 2
						break
					}
				}
			}
		}
	}
	return length
}

func dfs1443(mp map[int][]int, path *[]int, n int, used map[int]int) {
	if n == 0 {
		return
	}
	for i := 0; i < len(mp[n]); i++ {
		t := mp[n][i]
		if used[t] == 0 {
			used[t] = 1
			*path = append(*path, t)
			dfs1443(mp, path, t, used)
			if (*path)[len(*path)-1] == 0 {
				return
			}
			*path = (*path)[:len(*path)-1]
		}
	}
}
