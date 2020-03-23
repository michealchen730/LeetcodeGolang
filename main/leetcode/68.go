package main

import (
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	var wordlist []string
	var res []string
	j, length := 0, 0
	for j < len(words) {
		if len(wordlist) == 0 {
			wordlist = append(wordlist, words[j])
			length += len(words[j])
		} else {
			if length+len(words[j])+len(wordlist) > maxWidth {
				//结束，处理wordlist
				handleWordList(wordlist, &res, length, maxWidth)
				length = len(words[j])
				wordlist = []string{words[j]}
			} else {
				wordlist = append(wordlist, words[j])
				length += len(words[j])
			}
		}
		j++
	}
	handleLastWordList(wordlist, &res, length, maxWidth)
	return res
}
func handleWordList(wl []string, res *[]string, length, maxWidth int) {
	var temp strings.Builder
	if len(wl) == 1 {
		temp.WriteString(wl[0])
		for i := 0; i < maxWidth-length; i++ {
			temp.WriteByte(' ')
		}
	} else {
		space := maxWidth - length
		avgspace := space / (len(wl) - 1)
		remain := space % (len(wl) - 1)
		temp.WriteString(wl[0])
		for i := 1; i < len(wl); i++ {
			for j := 0; j < avgspace; j++ {
				temp.WriteByte(' ')
			}
			if i-1 < remain {
				temp.WriteByte(' ')
			}
			temp.WriteString(wl[i])
		}
	}
	*res = append(*res, temp.String())
}

func handleLastWordList(wl []string, res *[]string, length, maxWidth int) {
	var temp strings.Builder
	temp.WriteString(wl[0])
	for i := 1; i < len(wl); i++ {
		temp.WriteByte(' ')
		temp.WriteString(wl[i])
	}
	remain := (len(wl)-1)*1 + length
	for i := 0; i < maxWidth-remain; i++ {
		temp.WriteByte(' ')
	}
	*res = append(*res, temp.String())
}
