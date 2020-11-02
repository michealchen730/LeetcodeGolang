package main

func increasingBST(root *TreeNode) *TreeNode {

	var dfs func(root *TreeNode) (*TreeNode, *TreeNode)

	dfs = func(root *TreeNode) (*TreeNode, *TreeNode) {
		if root != nil {
			if root.Left == nil && root.Right == nil {
				return root, root
			}
			if root.Left != nil && root.Right != nil {
				n1, n2 := dfs(root.Left)
				n3, n4 := dfs(root.Right)
				root.Left = nil
				n2.Right = root
				root.Right = n3
				return n1, n4
			}
			if root.Left != nil && root.Right == nil {
				n1, n2 := dfs(root.Left)
				root.Left = nil
				n2.Right = root
				return n1, root
			}
			if root.Left == nil && root.Right != nil {
				n3, n4 := dfs(root.Right)
				root.Right = n3
				return root, n4
			}
		}
		return nil, nil
	}

	ans, _ := dfs(root)

	return ans

}

func increasingBST2(root *TreeNode) *TreeNode {
	var dfs func(root *TreeNode) (*TreeNode, *TreeNode)
	dfs = func(root *TreeNode) (*TreeNode, *TreeNode) {
		if root != nil {
			n1, n2 := dfs(root.Left)
			n3, n4 := dfs(root.Right)
			root.Left = nil
			if n2 != nil {
				n2.Right = root
			}
			root.Right = n3
			var ans1 *TreeNode
			if n1 != nil {
				ans1 = n1
			} else {
				ans1 = root
			}
			var ans2 *TreeNode
			if n4 != nil {
				ans2 = n4
			} else {
				ans2 = root
			}
			return ans1, ans2
		}
		return nil, nil
	}

	ans, _ := dfs(root)

	return ans

}
