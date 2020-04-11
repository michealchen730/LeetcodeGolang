package main

func constructFromPrePost(pre []int, post []int) *TreeNode {
	node := &TreeNode{Val: pre[0]}
	if len(pre) > 1 {
		i := 0
		for i < len(post) {
			if post[i] == pre[1] {
				break
			}
			i++
		}
		node.Left = constructFromPrePost(pre[1:i+2], post[0:i+1])
		if i+2 < len(pre) {
			node.Right = constructFromPrePost(pre[i+2:], post[i+1:len(post)-1])
		}
	}
	return node
}
