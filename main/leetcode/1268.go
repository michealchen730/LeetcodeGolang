package main

import (
	"strings"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	var res [][]string
	trie := Trie1268{}
	for _, v := range products {
		trie.Insert1268(v)
	}
	for i := 0; i < len(searchWord); i++ {
		temp := trie.StartsWith1268(searchWord[:i+1])
		var words []string
		if temp != nil {
			bytes := []byte(searchWord[:i+1])
			dfs1268(temp, &words, &bytes)
		}
		res = append(res, words)
	}
	return res
}

type Trie1268 struct {
	ending bool
	next   [26]*Trie1268
}

/** Inserts a word into the trie. */
func (this *Trie1268) Insert1268(word string) {
	temp := this
	for _, v := range word {
		nxt := v - 'a'
		if temp.next[nxt] == nil {
			temp.next[nxt] = &Trie1268{}
		}
		temp = temp.next[nxt]
	}
	temp.ending = true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie1268) StartsWith1268(prefix string) *Trie1268 {
	temp := this
	for _, v := range prefix {
		nxt := v - 'a'
		if temp.next[nxt] == nil {
			return nil
		} else {
			temp = temp.next[nxt]
		}
	}
	return temp
}

func dfs1268(trie *Trie1268, words *[]string, bytes *[]byte) {
	if len(*words) == 3 {
		return
	}
	if trie.ending == true {
		var sb strings.Builder
		sb.Write(*bytes)
		*words = append(*words, sb.String())
	}
	for i := 0; i < 26; i++ {
		if trie.next[i] != nil {
			*bytes = append(*bytes, byte('a'+i))
			dfs1268(trie.next[i], words, bytes)
			*bytes = (*bytes)[:len(*bytes)-1]
		}
	}
}
