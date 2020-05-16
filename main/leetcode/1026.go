package main

import "sort"

func maxAncestorDiff(root *TreeNode) int {
	var path []int
	mx := []int{0}
	dfs1026(root, &path, mx)
	return mx[0]
}

func dfs1026(root *TreeNode, path *[]int, mx []int) {
	*path = append(*path, root.Val)
	if root.Left == nil && root.Right == nil {
		cpy := make([]int, len(*path))
		copy(cpy, *path)
		sort.Ints(cpy)
		t := cpy[len(cpy)-1] - cpy[0]
		if t > mx[0] {
			mx[0] = t
		}
		*path = (*path)[:len(*path)-1]
		return
	}
	if root.Left != nil {
		dfs1026(root.Left, path, mx)
	}
	if root.Right != nil {
		dfs1026(root.Right, path, mx)
	}
	*path = (*path)[:len(*path)-1]
}

//错误的想法（觉得很有意义，因此保留了下来）
//func maxAncestorDiff(root *TreeNode) int {
//	mx:=[]int{100001,-1,0}
//	dfs1026(root,mx)
//	return mx[2]
//}
//func dfs1026(root *TreeNode, mx []int){
//	if root.Left==nil&&root.Right==nil{
//		mx[0]=min(root.Val,mx[0])
//		mx[1]=max(root.Val,mx[1])
//		return
//	}
//	if root.Left!=nil{
//		dfs1026(root.Left,mx)
//	}
//	if root.Right!=nil{
//		dfs1026(root.Right,mx)
//	}
//	t:=max(abs(mx[0]-root.Val),abs(mx[1]-root.Val))
//	if t>mx[2]{
//		mx[2]=t
//	}
//}
