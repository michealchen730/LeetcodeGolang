package main

func sortedArrayToBST(nums []int) *TreeNode {
	var buildBST func(nums []int, start, end int) *TreeNode
	buildBST = func(nums []int, start, end int) *TreeNode {
		if start > end {
			return nil
		}
		flag := start + (end-start)/2
		node := &TreeNode{Val: nums[flag]}
		node.Left = buildBST(nums, start, flag-1)
		node.Right = buildBST(nums, flag+1, end)
		return node
	}
	return buildBST(nums, 0, len(nums))
}
