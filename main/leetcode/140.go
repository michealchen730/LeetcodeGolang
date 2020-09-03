package main

//这题是我完全看不懂的题，只能先翻译题解，再看
//超时
//需要判断能够拼出来
func wordBreak140(s string, wordDict []string) []string {
	//wordDict用哈希表存一下
	mp := make(map[string]bool)
	for _, v := range wordDict {
		mp[v] = true
	}

	canbreak := make([]bool, len(s)+1)
	canbreak[0] = true
	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if mp[s[j:i]] && canbreak[j] {
				canbreak[i] = true
			}
		}
	}
	if !canbreak[len(s)] {
		return nil
	}

	//dp[i]表示s[:i]可被拆分成的句子
	dp := make([][]string, len(s)+1)
	initial := []string{""}
	dp[0] = initial

	for i := 1; i < len(s)+1; i++ {
		var temp []string
		for j := 0; j < i; j++ {
			if len(dp[j]) > 0 && mp[s[j:i]] {
				for _, v := range dp[j] {
					if v == "" {
						temp = append(temp, v+s[j:i])
					} else {
						temp = append(temp, v+" "+s[j:i])
					}
				}
			}
		}
		dp[i] = temp
	}
	return dp[len(s)]
}
