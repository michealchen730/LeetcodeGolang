package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var flag []*TreeNode
	if root != nil {
		if root == p || root == q {
			return root
		}
		mp := make(map[*TreeNode]*TreeNode)
		stack := []*TreeNode{root}
		for len(stack) != 0 {
			node := stack[0]
			if node.Left != nil {
				mp[node.Left] = node
				if node.Left == p || node.Left == q {
					flag = append(flag, node.Left)
				}
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				mp[node.Right] = node
				if node.Right == p || node.Right == q {
					flag = append(flag, node.Right)
				}
				stack = append(stack, node.Right)
			}
			if len(flag) == 2 {
				break
			}
			stack = stack[1:]
		}
		// 找到两个点了
		if len(flag) == 2 {
			mp2 := make(map[*TreeNode]int)
			for flag[0] != root {
				mp2[flag[0]]++
				flag[0] = mp[flag[0]]
			}
			mp2[root] = 1
			for mp2[flag[1]] == 0 {
				flag[1] = mp[flag[1]]
			}
			return flag[1]
		}
	}
	return nil
}
