package main

func recoverTree(root *TreeNode) {
	//还是应该理解为找逆序对，使用中序遍历
	var pre, x, y *TreeNode
	for root != nil {
		if root.Left != nil {
			node := root.Left
			node = node.Left
			for node.Right != nil && node.Right != root {
				node = node.Right
			}
			if node.Right == nil {
				node.Right = root
				root = root.Left
			} else {
				if pre != nil && root.Val < pre.Val {
					y = root
					if x == nil {
						x = pre
					}
				}
				pre = root
				node.Right = nil
				root = root.Right
			}
		} else {
			if pre != nil && root.Val < pre.Val {
				y = root
				if x == nil {
					x = pre
				}
			}
			pre = root
			root = root.Right
		}
	}
	x.Val, y.Val = y.Val, x.Val
}

func inOrder(root *TreeNode, nodes *[]*TreeNode, vals *[]int) {
	var stack []*TreeNode
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) != 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			*nodes = append(*nodes, root)
			*vals = append(*vals, root.Val)
			root = root.Right
		}
	}
}

//morris算法实现O(1)空间复杂度中序遍历
func morris(root *TreeNode) []*TreeNode {
	var res []*TreeNode
	for root != nil {
		if root.Left != nil {
			res = append(res, root)
			root = root.Right
		} else {
			pre := morrisGetPredecessor(root)
			if pre.Right == nil {
				pre.Right = root
				root = root.Left
			} else if pre.Right == root {
				pre.Right = nil
				res = append(res, root)
				root = root.Right
			}
		}
	}
	return res
}
func morrisVal(root *TreeNode) []int {
	var res []int
	for root != nil {
		if root.Left != nil {
			res = append(res, root.Val)
			root = root.Right
		} else {
			pre := morrisGetPredecessor(root)
			if pre.Right == nil {
				pre.Right = root
				root = root.Left
			} else if pre.Right == root {
				pre.Right = nil
				res = append(res, root.Val)
				root = root.Right
			}
		}
	}
	return res
}

func morrisGetPredecessor(root *TreeNode) *TreeNode {
	node := root
	if root.Left != nil {
		root = root.Left
		for root.Right != nil && root.Right != node {
			root = root.Right
		}
	}
	return root
}
