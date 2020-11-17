package main

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	//先决条件
	if voyage[0] != root.Val {
		return []int{-1}
	}
	//root.left==nil&&root.Right==nil
	if len(voyage) == 1 {
		return []int{0}
	}
	if root.Left == nil {
		return flipMatchVoyage(root.Right, voyage[1:])
	}
	if root.Right == nil {
		return flipMatchVoyage(root.Left, voyage[1:])
	}
	v1, v2 := -1, -1
	//根据先序遍历的切分数组
	for i := 0; i < len(voyage); i++ {
		if voyage[i] == root.Left.Val {
			v1 = i
		}
		if voyage[i] == root.Right.Val {
			v2 = i
		}
	}
	if v1 == -1 || v2 == -1 {
		return []int{-1}
	}
	var res []int
	var r1, r2 []int
	if v1 != 1 && v2 != 1 {
		return []int{-1}
	}
	if v1 != 1 {
		res = append(res, root.Val)
		r1, r2 = flipMatchVoyage(root.Left, voyage[v1:]), flipMatchVoyage(root.Right, voyage[v2:v1])
	} else {
		r1, r2 = flipMatchVoyage(root.Left, voyage[v1:v2]), flipMatchVoyage(root.Right, voyage[v2:])
	}
	if (len(r1) > 0 && r1[0] == -1) || (len(r2) > 0 && r2[0] == -1) {
		return []int{-1}
	} else {
		res = append(res, r1...)
		res = append(res, r2...)
		return res
	}
}
