package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var stack []int
	mp := make(map[int]int)
	for _, v := range nums2 {
		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] < v {
			mp[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}
	res := make([]int, len(nums1))
	for k, v := range nums1 {
		if ans, ok := mp[v]; !ok {
			res[k] = -1
		} else {
			res[k] = ans
		}
	}
	return res
}
