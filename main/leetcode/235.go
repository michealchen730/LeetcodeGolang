package main

func lowestCommonAncestor235(root, p, q *TreeNode) *TreeNode {
	//在二叉树的一左一右
	if max(p.Val, q.Val) >= root.Val && min(p.Val, q.Val) <= root.Val {
		return root
	}
	if max(p.Val, q.Val) < root.Val {
		return lowestCommonAncestor235(root.Left, p, q)
	}
	return lowestCommonAncestor235(root.Right, p, q)
}
