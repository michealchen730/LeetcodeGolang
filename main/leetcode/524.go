package main

func findLongestWord(s string, d []string) string {
	longest := 0
	var word string
	for _, v := range d {
		if isSubsequence(v, s) {
			if len(v) > longest {
				longest = len(v)
				word = v
			} else if len(v) == longest {
				for i := 0; i < len(v); i++ {
					if v[i] < word[i] {
						word = v
						break
					} else if v[i] > word[i] {
						break
					}
				}
			}
		}
	}
	return word
}
