package main

//这题简单来说，考以下哈希
//难一点，考一下字典树这种结构
//type WordsFrequency struct {
//	mp map[string]int
//}
//
//
//func Constructor(book []string) WordsFrequency {
//	res:=WordsFrequency{mp:make(map[string]int)}
//	for _,v:=range book{
//		res.mp[v]++
//	}
//	return res
//}
//
//
//func (this *WordsFrequency) Get(word string) int {
//	return this.mp[word]
//}

type WordsFrequency struct {
	ending int
	next   [26]*WordsFrequency
}

func Constructor1602(book []string) WordsFrequency {
	res := WordsFrequency{}
	for _, v := range book {
		res.Insert(v)
	}
	return res
}

func (this *WordsFrequency) Get(word string) int {
	temp := this
	for _, v := range word {
		nxt := v - 'a'
		if temp.next[nxt] == nil {
			return 0
		} else {
			temp = temp.next[nxt]
		}
	}
	return temp.ending
}

/** Inserts a word into the trie. */
func (this *WordsFrequency) Insert(word string) {
	temp := this
	for _, v := range word {
		nxt := v - 'a'
		if temp.next[nxt] == nil {
			temp.next[nxt] = &WordsFrequency{}
		}
		temp = temp.next[nxt]
	}
	temp.ending += 1
}
