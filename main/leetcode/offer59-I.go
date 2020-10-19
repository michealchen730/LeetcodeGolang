package main

func maxSlidingWindow3(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}

	//单调栈，存下标
	var stack []int

	for i := 0; i < k; i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}
		for len(stack) != 0 {
			if nums[stack[len(stack)-1]] <= nums[i] {
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, i)
	}
	res := []int{nums[stack[0]]}

	for i := 1; i < len(nums)-k+1; i++ {
		if stack[0] == i-1 {
			stack = stack[1:]
		}
		for len(stack) != 0 {
			if nums[stack[len(stack)-1]] <= nums[i+k-1] {
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, i+k-1)
		res = append(res, nums[stack[0]])
	}
	return res

}
