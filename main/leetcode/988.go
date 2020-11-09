package main

import (
	"strings"
)

func smallestFromLeaf(root *TreeNode) string {
	//第一个map存放所有的子节点-父节点
	mp := make(map[*TreeNode]*TreeNode)
	//数组存放所有的最小叶子节点
	mp2 := make(map[*TreeNode]int)

	//结果
	var res []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			if node.Left != nil {
				mp[node.Left] = node
				dfs(node.Left)
			}
			if node.Right != nil {
				mp[node.Right] = node
				dfs(node.Right)
			}
			if node.Left == nil && node.Right == nil {
				mp2[node] = node.Val
			}
		}
	}
	dfs(root)
FI:
	for len(mp2) != 0 {
		mn := 27
		for _, v := range mp2 {
			mn = min(mn, v)
		}
		res = append(res, mn)
		mp3 := make(map[*TreeNode]int)
		for k, v := range mp2 {
			if v == mn {
				if _, ok := mp[k]; ok {
					mp3[mp[k]] = mp[k].Val
				} else {
					break FI
				}
			}
		}
		mp2 = mp3
	}

	var r strings.Builder

	for _, v := range res {
		r.WriteByte(byte('a' + v))
	}
	return r.String()
}
