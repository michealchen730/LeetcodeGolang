package main

type WordDictionary struct {
	trie *Trie
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{trie: &Trie{}}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	this.trie.Insert(word)
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.trie.Search211(word)
}

func (this *Trie) Search211(word string) bool {
	temp := this
	for k, v := range word {
		if v == '.' {
			if k == len(word)-1 {
				for _, v1 := range temp.next {
					if v1 != nil && v1.ending {
						return true
					}
				}
				return false
			} else {
				for _, v1 := range temp.next {
					if v1 != nil && v1.Search211(word[k+1:]) {
						return true
					}
				}
				return false
			}
		} else {
			nxt := v - 'a'
			if temp.next[nxt] == nil {
				return false
			} else {
				temp = temp.next[nxt]
			}
		}
	}
	return temp.ending
}
