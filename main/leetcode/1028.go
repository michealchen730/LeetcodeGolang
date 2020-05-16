package main

import (
	"strconv"
)

func recoverFromPreorder(S string) *TreeNode {
	stack := []int{0}
	var nodes []*TreeNode
	for i := 0; i < len(S); i++ {
		if S[i] == '-' {
			j := i
			for i < len(S) && S[i] == '-' {
				i++
			}
			depth := i - j
			stack = append(stack, depth)
			i--
			continue
		}
		if S[i] != '-' {
			j := i
			for i < len(S) && S[i] != '-' {
				i++
			}
			t, _ := strconv.Atoi(S[j:i])
			//做一定的处理
			node := &TreeNode{Val: t}
			nodes = append(nodes, node)
			if len(stack) != 1 {
				for d := len(stack) - 1; d >= 0; d-- {
					if stack[d] < stack[len(stack)-1] {
						if d == len(stack)-2 {
							nodes[d].Left = node
						} else {
							nodes[d].Right = node
						}
						break
					}
				}
			}
			i--
		}
	}
	return nodes[0]
}
