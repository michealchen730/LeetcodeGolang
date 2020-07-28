package main

func convertBiNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	//先处理右子树
	root.Right = convertBiNode(root.Right)
	//再处理左子树
	//如果左子树不为空
	//那么返回的头节点不再是当前root节点
	if root.Left != nil {
		res := convertBiNode(root.Left)
		node := res
		for node.Right != nil {
			node = node.Right
		}
		node.Right = root
		root.Left = nil
		return res
	} else {
		return root
	}
}
