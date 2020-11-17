package main

import "sort"

func verticalTraversal(root *TreeNode) [][]int {
	mp := make(map[int][]int)

	queue := []*TreeNode{root}
	val := []int{0}

	for len(queue) != 0 {
		l := len(queue)
		mp2 := make(map[int][]int)
		for i := 0; i < l; i++ {
			n, p := queue[i], val[i]
			mp2[p] = append(mp2[p], n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
				val = append(val, p-1)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
				val = append(val, p+1)
			}
		}
		queue = queue[l:]
		val = val[l:]
		for k, _ := range mp2 {
			sort.Ints(mp2[k])
			mp[k] = append(mp[k], mp2[k]...)
		}
	}

	var arr []int
	for k, _ := range mp {
		arr = append(arr, k)
	}
	sort.Ints(arr)

	var res [][]int
	for _, v := range arr {
		res = append(res, mp[v])
	}
	return res
}
