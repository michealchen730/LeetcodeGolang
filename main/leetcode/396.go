package main

func maxRotateFunction(nums []int) int {
	sum, f := 0, 0
	for k, v := range nums {
		sum += v
		f += k * v
	}

	res := f

	for t := len(nums) - 1; t >= 0; t-- {
		f = f - nums[t]*len(nums) + sum
		res = max(res, f)
	}
	return res
}
