package main

import (
	"math"
	"strings"
)

//"abcacb"
//"acbabc"

//这题对于贪心原则的使用需要证明

//但思路是这样的
//找到确保所有字母都出现一次的最后的位置，比如"abcacb"中的第二个a，就是可以确保所有字母都出现一次的最后的位置
//然后从头开始，找字典序最小的字母，一直找到那个最后的位置为止，比如"abcacb"，就是从头开始一直找到第二个"a"的位置停下
//为什么从头开始，比如在"abcacb"中，我们就需要第一个a而不是第二个a（这也是需要证明正确性的一点）
//然后对字符串做如下处理，分割字符串，从找到的那个最小字典序字母那里开始分割，然后替换所有最小字典序字母为""

func removeDuplicateLetters(s string) string {
	if len(s) == 0 {
		return ""
	}
	//存下每个字母出现的最先位置和最后位置，这样可以找到那个确保所有字母都出现一次的最靠后的位置
	lst := make([]int, 26)
	frs := make([]int, 26)
	for k, _ := range s {
		lst[s[k]-'a'] = k + 1
		if frs[s[k]-'a'] == 0 {
			frs[s[k]-'a'] = k + 1
		}
	}

	var res strings.Builder
	loc := math.MaxInt32

	//找到那个确保所有字母都出现一次的最靠后的位置
	for _, v := range lst {
		if v != 0 && v < loc {
			loc = v
		}
	}

	mn := 0
	mnloc := 0
	//从头开始，找到最小字典序的字母
	for k, v := range frs {
		if v != 0 && v <= loc { //必然存在
			mn = k
			mnloc = v
			break
		}
	}

	//记录最小字典序字母
	res.WriteByte(byte('a' + mn))

	//分割
	rmn := s[mnloc:]
	//替换
	newstr := strings.ReplaceAll(rmn, string(byte('a'+mn)), "")

	//递归
	res.WriteString(removeDuplicateLetters(newstr))
	return res.String()
}

func removeDuplicateLetters2(s string) string {
	var stack []byte

	arr := make([]int, 26)

	for k, _ := range s {
		arr[s[k]-'a'] = k + 1
	}

	mp := make(map[byte]int)

	for k, _ := range s {
		//栈中有重复元素，那么直接跳过后续步骤
		if mp[s[k]] != 0 {
			continue
		}

		//判断栈顶元素出栈
		for len(stack) > 0 {
			if stack[len(stack)-1] > s[k] && arr[stack[len(stack)-1]-'a'] > k {
				mp[stack[len(stack)-1]]--
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}

		stack = append(stack, s[k])
		mp[s[k]]++
	}

	return string(stack)
}
