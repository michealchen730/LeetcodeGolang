package main

func findSubstring(s string, words []string) []int {
	var res []int
	if len(words) == 0 {
		return res
	}
	length := len(words[0])
	mp := make(map[string]int)
	for i := 0; i < len(words); i++ {
		mp[words[i]]++
	}
	trie := Trie{}
	for k, _ := range mp {
		trie.Insert(k)
	}
	for i := 0; i <= len(s)-length*len(words); i++ {
		if trie.Search(s[i : i+length]) {
			mp2 := make(map[string]int)
			mp2[s[i:i+length]]++
			flag := 0
			for j := i + length; j < i+len(words)*length; j += length {
				if trie.Search(s[j : j+length]) {
					mp2[s[j:j+length]]++
					if mp2[s[j:j+length]] > mp[s[j:j+length]] {
						flag = 1
						break
					}
				} else {
					flag = 1
					break
				}
			}
			if flag == 0 {
				res = append(res, i)
			}
		}
	}
	return res
}
