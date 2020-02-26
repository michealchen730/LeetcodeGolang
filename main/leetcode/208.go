package main

//type Trie struct {
//	ending bool
//	next [26]*Trie
//}
//
//
///** Initialize your data structure here. */
//func Constructor() Trie {
//	return Trie{}
//}
//
//
///** Inserts a word into the trie. */
//func (this *Trie) Insert(word string)  {
//	temp:=this
//	for _,v:=range word {
//		nxt:=v-'a'
//		if temp.next[nxt]==nil {
//			temp.next[nxt]=&Trie{}
//		}
//		temp=temp.next[nxt]
//	}
//	temp.ending=true
//}
//
//
///** Returns if the word is in the trie. */
//func (this *Trie) Search(word string) bool {
//	temp:=this
//	for _,v:=range word{
//		nxt:=v-'a'
//		if temp.next[nxt]==nil {
//			return false
//		}else {
//			temp=temp.next[nxt]
//		}
//	}
//	return temp.ending
//}
//
//
///** Returns if there is any word in the trie that starts with the given prefix. */
//func (this *Trie) StartsWith(prefix string) bool {
//	temp:=this
//	for _,v:=range prefix{
//		nxt:=v-'a'
//		if temp.next[nxt]==nil {
//			return false
//		}else{
//			temp=temp.next[nxt]
//		}
//	}
//	return true
//}
