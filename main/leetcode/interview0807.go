package main

func permutation(S string) []string {
	if len(S) == 0 {
		return []string{}
	}
	var res []string
	var path []byte
	used := make([]bool, len(S))
	dfs0807(S, 0, used, &path, &res)
	return res

}

func dfs0807(s string, n int, used []bool, path *[]byte, res *[]string) {
	if n == len(s) {
		*res = append(*res, string(*path))
	}
	for i := 0; i < len(used); i++ {
		if !used[i] {
			used[i] = true
			*path = append(*path, s[i])
			dfs0807(s, n+1, used, path, res)
			*path = (*path)[:len(*path)-1]
			used[i] = false
		}
	}
}
