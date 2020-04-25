package main

import "math"

func getHappyString(n int, k int) string {
	sum := 3 * int(math.Pow(2, float64(n-1)))
	if k > sum {
		return ""
	}
	var res []string
	var str []byte
	buildHappyString(n, k, &str, &res)
	return res[k-1]
}

func buildHappyString(n int, k int, word *[]byte, str *[]string) {
	if len(*str) == k {
		return
	}
	if len(*word) == n {
		*str = append(*str, string(*word))
		return
	}
	if len(*word) == 0 {
		*word = append(*word, 'a')
		buildHappyString(n, k, word, str)
		*word = (*word)[:len(*word)-1]
		*word = append(*word, 'b')
		buildHappyString(n, k, word, str)
		*word = (*word)[:len(*word)-1]
		*word = append(*word, 'c')
		buildHappyString(n, k, word, str)
		*word = (*word)[:len(*word)-1]
	} else {
		temp := (*word)[len(*word)-1]
		if temp == 'a' {
			*word = append(*word, 'b')
			buildHappyString(n, k, word, str)
			*word = (*word)[:len(*word)-1]
			*word = append(*word, 'c')
			buildHappyString(n, k, word, str)
			*word = (*word)[:len(*word)-1]
		} else if temp == 'b' {
			*word = append(*word, 'a')
			buildHappyString(n, k, word, str)
			*word = (*word)[:len(*word)-1]
			*word = append(*word, 'c')
			buildHappyString(n, k, word, str)
			*word = (*word)[:len(*word)-1]
		} else {
			*word = append(*word, 'a')
			buildHappyString(n, k, word, str)
			*word = (*word)[:len(*word)-1]
			*word = append(*word, 'b')
			buildHappyString(n, k, word, str)
			*word = (*word)[:len(*word)-1]
		}
	}
}
