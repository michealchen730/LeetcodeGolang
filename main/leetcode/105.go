package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	node := &TreeNode{Val: preorder[0]}
	i := 0
	//分情况讨论
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	//存在右子树
	if i != len(inorder)-1 {
		node.Right = buildTree(preorder[i+1:], inorder[i+1:])
	}
	//存在左子树
	if i > 0 {
		node.Left = buildTree(preorder[1:i+1], inorder[0:i])
	}
	return node
}
