package main

type FindElements struct {
	newTree *TreeNode
}

func Constructor1261(root *TreeNode) FindElements {
	queue := []*TreeNode{root}
	root.Val = 0
	for len(queue) != 0 {
		temp := queue[0]
		if temp.Left != nil {
			temp.Left.Val = 2*temp.Val + 1
			queue = append(queue, temp.Left)
		}
		if temp.Right != nil {
			temp.Right.Val = 2*temp.Val + 2
			queue = append(queue, temp.Right)
		}
		queue = queue[1:]
	}
	return FindElements{newTree: root}
}

func (this *FindElements) Find(target int) bool {
	if target == 0 {
		return true
	}
	arr := []int{0}
	sum := 1
	for arr[len(arr)-1] < target {
		sum *= 2
		arr = append(arr, arr[len(arr)-1]+sum)
	}
	temp := arr[len(arr)-1] - arr[len(arr)-2]
	pos := arr[len(arr)-2] + 1
	tempnode := this.newTree
	for temp != 1 {
		temp = temp / 2
		if target >= temp+pos {
			tempnode = tempnode.Right
			pos = pos + temp
		} else {
			tempnode = tempnode.Left
		}
		if tempnode == nil {
			return false
		}
	}
	return true
}
