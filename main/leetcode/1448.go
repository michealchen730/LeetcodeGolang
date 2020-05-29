package main

func goodNodes(root *TreeNode) int {
	var path []int
	res := []int{0}
	if root != nil {
		dfs0516(root, &path, res)
	}
	return res[0]
}

func dfs0516(root *TreeNode, path *[]int, arr []int) {
	//应该有更好的办法，记录下当前最大值的这么一个状态，然后可以回溯回去，实际上，我们可以改为记录节点下标
	*path = append(*path, root.Val)
	flag := true
	for _, v := range *path {
		if v > (*path)[len(*path)-1] {
			flag = false
			break
		}
	}
	if flag {
		arr[0]++
	}
	if root.Left == nil && root.Right == nil {
		*path = (*path)[:len(*path)-1]
		return
	}
	if root.Left != nil {
		dfs0516(root.Left, path, arr)
	}
	if root.Right != nil {
		dfs0516(root.Right, path, arr)
	}
	*path = (*path)[:len(*path)-1]
}
