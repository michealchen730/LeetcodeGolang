package main

func findAndReplacePattern(words []string, pattern string) []string {
	var res []string
	for _, v := range words {
		flag := 0
		mp1 := make(map[byte]byte)
		mp2 := make(map[byte]byte)
		for i := 0; i < len(v); i++ {
			b1, ok1 := mp1[v[i]]
			b2, ok2 := mp2[pattern[i]]
			if (!ok1) && (!ok2) {
				mp1[v[i]] = pattern[i]
				mp2[pattern[i]] = v[i]
			}
			if (ok1 && !ok2) || (!ok1 && ok2) || (ok1 && ok2 && (b1 != pattern[i] || b2 != v[i])) {
				flag = 1
				break
			}
		}
		if flag == 0 {
			res = append(res, v)
		}
	}
	return res
}
