package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if q == nil && p == nil {
		return true
	}
	if p != nil && q != nil && q.Val == p.Val {
		return isSameTree(q.Left, p.Left) && isSameTree(q.Right, p.Right)
	} else {
		return false
	}
}
