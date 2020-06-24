package main

func wordBreak(s string, wordDict []string) bool {
	mp := make(map[string]int)
	for _, v := range wordDict {
		mp[v]++
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(dp); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] == true && mp[s[j:i]] != 0 {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
