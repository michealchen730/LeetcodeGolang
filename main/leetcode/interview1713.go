package main

func respace(dictionary []string, sentence string) int {
	trie := Trie{}
	for _, v := range dictionary {
		trie.Insert(v)
	}
	dp := make([]int, len(sentence)+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = dp[i-1] + 1
		for j := i - 1; j >= 0; j-- {
			if trie.Search(sentence[j:i]) {
				dp[i] = min(dp[i], dp[j])
			}
		}
	}
	return dp[len(sentence)]
}
