package main

type MagicDictionary struct {
	mp map[int][]string
}

/** Initialize your data structure here. */
func Constructor676() MagicDictionary {
	return MagicDictionary{mp: make(map[int][]string)}
}

/** Build a dictionary through a list of words */
func (this *MagicDictionary) BuildDict(dict []string) {
	for _, v := range dict {
		this.mp[len(v)] = append(this.mp[len(v)], v)
	}
}

/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (this *MagicDictionary) Search(word string) bool {
	words := this.mp[len(word)]
	for _, v := range words {
		t := 0
		for i := 0; i < len(word); i++ {
			if v[i] != word[i] {
				t++
			}
			if t > 1 {
				break
			}
		}
		if t == 1 {
			return true
		}
	}
	return false
}
