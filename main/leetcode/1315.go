package main

func sumEvenGrandparent(root *TreeNode) int {
	ans := 0

	var dfs1315 func(root *TreeNode, grandparent int)
	dfs1315 = func(root *TreeNode, parentValue int) {
		if (root.Left != nil || root.Right != nil) && parentValue%2 == 0 {
			if root.Left != nil {
				ans += root.Left.Val
			}
			if root.Right != nil {
				ans += root.Right.Val
			}
		}
		if root.Left != nil {
			dfs1315(root.Left, root.Val)
		}
		if root.Right != nil {
			dfs1315(root.Right, root.Val)
		}
	}

	dfs1315(root, 1)
	return ans
}

//理解错了题意，这个的解法是把所有偶数的祖父节点加起来
//懒得删了
func sumEvenGrandparentError(root *TreeNode) int {
	ans := 0

	var getDepth func(root *TreeNode) int
	getDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		res := max(getDepth(root.Left), getDepth(root.Right)) + 1
		if res > 1 && root.Val%2 == 0 {
			ans += root.Val
		}
		return res
	}

	getDepth(root)
	return ans
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}
