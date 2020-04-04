package main

import (
	"regexp"
	"strconv"
	"strings"
)

//考虑递归
func decodeString(s string) string {
	var res strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			res.WriteByte(s[i])
		} else {
			j := i + 1
			for j < len(s) {
				if s[j] == '[' {
					break
				}
				j++
			}
			flag := 1
			k := j + 1
			for k < len(s) {
				if s[k] == '[' {
					flag++
				}
				if s[k] == ']' {
					flag--
					if flag == 0 {
						break
					}
				}
				k++
			}
			nums, _ := strconv.Atoi(s[i:j])
			str := decodeString(s[j+1 : k])
			for m := 0; m < nums; m++ {
				res.WriteString(str)
			}
			i = k
		}
	}
	return res.String()
}

//考虑正则表达式实现
func decodeString2(s string) string {
	res := regexp.MustCompile("\\d+\\[\\w+]")
	res2 := regexp.MustCompile("\\d+")
	res3 := regexp.MustCompile("\\[\\w+]")
	for res.MatchString(s) {
		s = res.ReplaceAllStringFunc(s, func(tes string) string {
			var temp strings.Builder
			nums, _ := strconv.Atoi(res2.FindAllString(tes, -1)[0])
			words := res3.FindAllString(tes, -1)[0]
			word := words[1 : len(words)-1]
			for i := 0; i < nums; i++ {
				temp.WriteString(word)
			}
			return temp.String()
		})
	}
	return s
}
