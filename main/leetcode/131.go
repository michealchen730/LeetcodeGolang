package main

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{{}}
	}
	dp := make([][]int, len(s))
	for k, _ := range dp {
		dp[k] = make([]int, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if i == j {
				dp[i][j] = 1
			}
			if j-i == 1 && s[i] == s[j] {
				dp[i][j] = 1
			}
			if j-i > 1 {
				if dp[i+1][j-1] == 1 && s[i] == s[j] {
					dp[i][j] = 1
				}
			}
		}
	}
	var res [][]string
	var path []string
	getPartition(0, 1, s, &path, &res, dp)
	return res
}

func getPartition(last int, i int, s string, path *[]string, res *[][]string, dp [][]int) {
	if i == len(s) {
		if dp[last][i-1] == 1 {
			*path = append(*path, s[last:])
			cpy := make([]string, len(*path))
			copy(cpy, *path)
			*res = append(*res, cpy)
			temp2 := *path
			temp2 = temp2[:len(temp2)-1]
			*path = temp2
		}
		return
	}
	getPartition(last, i+1, s, path, res, dp)
	if dp[last][i-1] == 1 {
		*path = append(*path, s[last:i])
		getPartition(i, i+1, s, path, res, dp)
		temp2 := *path
		temp2 = temp2[:len(temp2)-1]
		*path = temp2
	}
}
