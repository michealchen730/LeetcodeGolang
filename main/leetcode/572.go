package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isSubtree(s *TreeNode, t *TreeNode) bool {
	s1 := preOrderDFS572(s)
	s2 := preOrderDFS572(t)
	fmt.Println(s1)
	fmt.Println(s2)
	return strStr(s1, s2) != -1
}

func preOrderDFS572(s *TreeNode) string {
	var stack []*TreeNode
	if s != nil {
		var res strings.Builder
		stack = append(stack, s)
		for len(stack) != 0 {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res.WriteByte(' ')
			res.WriteString(strconv.Itoa(temp.Val))
			if temp.Right != nil {
				stack = append(stack, temp.Right)
			}
			if temp.Left != nil {
				stack = append(stack, temp.Left)
			} else {
				res.WriteByte('l')
			}
			if temp.Right == nil {
				res.WriteByte('r')
			}
		}
		return res.String()
	}
	return ""
}

func getNext(s string) []int {
	next := make([]int, len(s))
	next[0] = -1
	i, j := 0, -1
	for i < len(s)-1 {
		if j == -1 || s[i] == s[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	i, j := 0, 0
	next := getNext(needle)
	for i < len(haystack) {
		if j == len(needle)-1 && needle[j] == haystack[i] {
			return i - j
		}
		if j == -1 || needle[j] == haystack[i] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	return -1
}
