package main

//递归写法
func postorderTraversal(root *TreeNode) []int {
	var res []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Left)
			dfs(root.Right)
			res = append(res, root.Val)
		}
	}

	dfs(root)
	return res
}

//非递归写法
//双栈法

func postorderTraversal2(root *TreeNode) []int {
	var stack1 []*TreeNode
	var res []int

	if root != nil {
		stack1 = append(stack1, root)
		for len(stack1) != 0 {
			node := stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
			res = append(res, node.Val)
			if node.Left != nil {
				stack1 = append(stack1, node.Left)
			}
			if node.Right != nil {
				stack1 = append(stack1, node.Right)
			}
		}
		i, j := 0, len(res)-1
		for i < j {
			res[i], res[j] = res[j], res[i]
			i++
			j--
		}
	}
	return res
}
