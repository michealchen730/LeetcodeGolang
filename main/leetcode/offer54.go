package main

func kthLargest(root *TreeNode, k int) int {
	t, res := 0, 0
	var traverseO54 func(root *TreeNode)
	traverseO54 = func(root *TreeNode) {
		if root != nil {
			if root.Right != nil {
				traverseO54(root.Right)
			}
			t += 1
			if t == k {
				res = root.Val
			}
			if t >= k {
				return
			}
			if root.Left != nil {
				traverseO54(root.Left)
			}
		}
	}
	traverseO54(root)
	return res
}
