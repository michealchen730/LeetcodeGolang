package main

import "strings"

func uniqueMorseRepresentations(words []string) int {
	arr := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	hashtable := make(map[string]int)
	for _, word := range words {
		var str strings.Builder
		for _, c := range word {
			str.WriteString(arr[int(c-'a')])
		}
		hashtable[str.String()]++
	}
	return len(hashtable)
}
