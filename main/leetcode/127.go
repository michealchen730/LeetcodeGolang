package main

import "fmt"

//广度优先搜索
func ladderLength(beginWord string, endWord string, wordList []string) int {
	l := len(beginWord)
	mp := make(map[string][]int)
	for k, v := range wordList {
		for i := 0; i < l; i++ {
			newword := v[:i] + "*" + v[i+1:]
			mp[newword] = append(mp[newword], k)
		}
	}

	used, res := make([]bool, len(wordList)), 0
	wordList = append(wordList, beginWord)
	queue := []int{len(wordList) - 1}
	for len(queue) != 0 {
		res++
		size := len(queue)

		for i := 0; i < size; i++ {
			if wordList[queue[i]] == endWord {
				return res
			}

			for k, _ := range wordList[i] {
				newword := wordList[queue[i]][:k] + "*" + wordList[queue[i]][k+1:]
				for _, v := range mp[newword] {
					if !used[v] {
						used[v] = true
						queue = append(queue, v)
					}
				}
			}
		}
		queue = queue[size:]
	}
	return 0
}

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}
