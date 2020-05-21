package main

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	dp := make([][]bool, len(s))
	start, end := 0, 1
	mx := 1
	for k, _ := range dp {
		dp[k] = make([]bool, len(s))
	}
	for k := len(s) - 1; k >= 0; k-- {
		for j := len(s) - 1; j >= k; j-- {
			if j == k {
				dp[k][j] = true
				continue
			}
			if s[j] == s[k] {
				if j-k == 1 {
					dp[k][j] = true
				} else {
					dp[k][j] = dp[k+1][j-1]
				}
			}
			if dp[k][j] && j-k+1 > mx {
				mx = j - k + 1
				start = k
				end = j + 1
			}
		}
	}
	return s[start:end]
}
