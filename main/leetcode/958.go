package main

func isCompleteTree(root *TreeNode) bool {
	hashtable := map[*TreeNode]int{root: 1}
	queue := []*TreeNode{root}
	mx := 1
	for len(queue) != 0 {
		node := queue[0]
		idx := hashtable[node]
		if node.Left != nil {
			hashtable[node.Left] = idx * 2
			queue = append(queue, node.Left)
			mx = idx * 2
		}
		if node.Right != nil {
			hashtable[node.Right] = idx*2 + 1
			queue = append(queue, node.Right)
			mx = idx*2 + 1
		}
		queue = queue[1:]
	}
	return mx == len(hashtable)
}
