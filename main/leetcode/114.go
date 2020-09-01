package main

func flatten(root *TreeNode) {

	var dfs func(root *TreeNode) *TreeNode

	//dfs拿到的是展开后最右下角的元素
	dfs = func(root *TreeNode) *TreeNode {
		if root != nil {
			t1 := dfs(root.Left)
			t2 := dfs(root.Right)
			r := root.Right
			root.Right = root.Left
			root.Left = nil
			t1.Right = r

			if t2 != nil {
				return t2
			} else {
				return t1
			}
		} else {
			return nil
		}
	}
	dfs(root)
}
