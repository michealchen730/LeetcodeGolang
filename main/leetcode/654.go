package main

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[mx] {
			mx = i
		}
	}
	root := &TreeNode{Val: nums[mx]}
	root.Left = constructMaximumBinaryTree(nums[:mx])
	root.Right = constructMaximumBinaryTree(nums[mx+1:])
	return root
}
