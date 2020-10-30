package main

func expressiveWords(S string, words []string) int {
	res := 0
	for _, v := range words {
		i, j := 0, 0
		for i < len(S) && j < len(v) {
			if S[i] != v[j] {
				break
			} else {
				m, n := i, j
				for m < len(S) && S[m] == S[i] {
					m++
				}
				for n < len(v) && v[n] == v[j] {
					n++
				}

				//如果长度达到3及以上，可以忽略n
				//如果没达到3，那么需要字母组长度一样
				//m-i>n-j是为了预防 "aaa" ["aaaa"]这种测试样例
				if (m-i > n-j && m-i > 2) || m-i == n-j {
					i, j = m, n
				} else {
					break
				}
			}
		}
		if i == len(S) && j == len(v) {
			res++
		}
	}
	return res
}
