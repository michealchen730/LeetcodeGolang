package main

import "sort"

//func verticalTraversal(root *TreeNode) [][]int {
//	mp := make(map[int][]int)
//
//	queue := []*TreeNode{root}
//	val := []int{0}
//
//	for len(queue) != 0 {
//		l := len(queue)
//		// 层次遍历的方式，把每一层的节点都收集起来
//		mp2 := make(map[int][]int)
//		for i := 0; i < l; i++ {
//			n, p := queue[i], val[i]
//			mp2[p] = append(mp2[p], n.Val)
//			if n.Left != nil {
//				queue = append(queue, n.Left)
//				val = append(val, p-1)
//			}
//			if n.Right != nil {
//				queue = append(queue, n.Right)
//				val = append(val, p+1)
//			}
//		}
//		queue = queue[l:]
//		val = val[l:]
//		// 对每个mp2都做了次排序
//		for k, _ := range mp2 {
//			sort.Ints(mp2[k])
//			mp[k] = append(mp[k], mp2[k]...)
//		}
//	}
//
//	var arr []int
//	for k, _ := range mp {
//		arr = append(arr, k)
//	}
//	sort.Ints(arr)
//
//	var res [][]int
//	for _, v := range arr {
//		res = append(res, mp[v])
//	}
//	return res
//}

func verticalTraversal(root *TreeNode) [][]int {
	mp := make(map[int][][]int)
	var dfs func(row, col int, r *TreeNode)
	dfs = func(row, col int, r *TreeNode) {
		if r == nil {
			return
		}
		mp[col] = append(mp[col], []int{row, r.Val})
		dfs(row+1, col-1, r.Left)
		dfs(row+1, col+1, r.Right)
	}
	dfs(0, 0, root)
	cols := make([]int, len(mp))
	tmp := 0
	for k, _ := range mp {
		cols[tmp] = k
		tmp++
	}
	sort.Ints(cols)
	var res [][]int
	for _, v := range cols {
		t := mp[v]
		sort.Slice(t, func(i, j int) bool {
			if t[i][0] < t[j][0] {
				return true
			} else if t[i][0] > t[j][0] {
				return false
			} else {
				return t[i][1] < t[j][1]
			}
		})
		arr := make([]int, len(t))
		for k, v2 := range t {
			arr[k] = v2[1]
		}
		res = append(res, arr)
	}
	return res
}
