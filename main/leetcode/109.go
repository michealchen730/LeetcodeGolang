package main

//由一年前的java代码改写过来
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	flag := 0
	node := head
	for node != nil {
		flag++
		node = node.Next
	}
	node2 := head
	nums := make([]int, flag)
	for i := 0; i < len(nums); i++ {
		nums[i] = node2.Val
		node2 = node2.Next
	}
	res := buildBST109(nums, 0, flag-1)
	return res
}

func buildBST109(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	node := &TreeNode{Val: nums[mid]}
	node.Left = buildBST109(nums, start, mid-1)
	node.Right = buildBST109(nums, mid+1, end)
	return node
}
