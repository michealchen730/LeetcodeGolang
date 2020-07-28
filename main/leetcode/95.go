package main

func generateTrees(n int) []*TreeNode {
	arr := make([]int, n)
	for i := 1; i <= n; i++ {
		arr[i-1] = i
	}
	res := generateTreesByArr(arr)
	return res
}

func generateTreesByArr(arr []int) []*TreeNode {
	if len(arr) == 0 {
		return []*TreeNode{}
	}
	if len(arr) == 1 {
		return []*TreeNode{&TreeNode{Val: arr[0]}}
	}
	var nodes []*TreeNode
	for i := 0; i < len(arr); i++ {
		left := generateTreesByArr(arr[:i])
		right := generateTreesByArr(arr[i+1:]) //这一句可能有问题
		m := max(1, len(left))
		n := max(1, len(right))
		for l := 0; l < m; l++ {
			for j := 0; j < n; j++ {
				node := &TreeNode{Val: arr[i]}
				if len(left) > 0 {
					node.Left = left[l]
				}
				if len(right) > 0 {
					node.Right = right[j]
				}
				nodes = append(nodes, node)
			}
		}
	}
	return nodes
}
