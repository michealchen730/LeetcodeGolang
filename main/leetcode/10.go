package main

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(p)+1)
	for k, _ := range dp {
		dp[k] = make([]bool, len(s)+1)
	}
	dp[0][0] = true
	for i := 1; i <= len(p); i++ {
		if p[i-1] != '*' {
			dp[i][0] = false
		} else {
			dp[i][0] = dp[i-2][0]
		}
	}
	for i := 1; i <= len(s); i++ {
		dp[0][i] = false
	}
	for i := 1; i <= len(p); i++ {
		for j := 1; j <= len(s); j++ {
			if i == 8 && j == 2 {
			}
			if p[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}
			if p[i-1] == '*' {
				if p[i-2] != s[j-1] {
					if p[i-2] != '.' {
						dp[i][j] = dp[i-2][j]
					} else {
						dp[i][j] = dp[i-2][j] || dp[i-1][j] || dp[i-1][j-1] || dp[i][j-1]
					}
				}
				if p[i-2] == s[j-1] {
					dp[i][j] = dp[i-1][j-1] || dp[i-1][j] || dp[i][j-1] || dp[i-2][j]
				}
				continue
			}
			if p[i-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
				continue
			}
			dp[i][j] = false
		}
	}
	return dp[len(p)][len(s)]
}
