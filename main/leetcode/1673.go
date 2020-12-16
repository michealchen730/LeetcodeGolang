package main

func mostCompetitive(nums []int, k int) []int {
	var stack []int

	for i := 0; i < len(nums); i++ {
		if len(stack) == 0 || len(nums)-i+len(stack) == k {
			stack = append(stack, nums[i])
			continue
		}
		for len(stack) > 0 && len(stack)+len(nums)-i > k && stack[len(stack)-1] > nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) < k {
			stack = append(stack, nums[i])
		}
	}
	return stack

}
