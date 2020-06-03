package main

func pseudoPalindromicPaths(root *TreeNode) int {
	var path []int
	arr := []int{0}
	dfs0524(root, &path, arr)
	return arr[0]
}

func dfs0524(root *TreeNode, path *[]int, arr []int) {
	*path = append(*path, root.Val)
	if root.Left == nil && root.Right == nil {
		if checkPali(*path) {
			arr[0]++
		}
		*path = (*path)[0 : len(*path)-1]
		return
	}
	if root.Left != nil {
		dfs0524(root.Left, path, arr)
	}
	if root.Right != nil {
		dfs0524(root.Right, path, arr)
	}
	*path = (*path)[0 : len(*path)-1]
}

func checkPali(arr []int) bool {
	a := make([]int, 10)
	for _, v := range arr {
		a[v]++
	}
	flag := 0
	for _, v := range a {
		if v%2 != 0 {
			flag++
		}
		if flag > 1 {
			return false
		}
	}
	return true
}
