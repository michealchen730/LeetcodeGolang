package main

//字典树的思路应该可行
func removeSubfolders(folder []string) []string {
	mp := make(map[string]bool)
	for _, v := range folder {
		mp[v] = true
	}
	isSub := make([]bool, len(folder))
	for k, v := range folder {
		for j := 0; j < len(v); j++ {
			if v[j] == '/' && j != 0 {
				if mp[v[:j]] {
					isSub[k] = true
					break
				}
			}
		}
	}

	var res []string
	for k, v := range isSub {
		if !v {
			res = append(res, folder[k])
		}
	}
	return res
}
