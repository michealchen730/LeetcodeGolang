package main

func maxProduct(nums []int) int {
	res := nums[0]
	minf, maxf := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		t1 := min(nums[i], nums[i]*minf)
		t1 = min(t1, nums[i]*maxf)
		t2 := max(nums[i], nums[i]*maxf)
		t2 = max(t2, nums[i]*minf)
		minf, maxf = t1, t2
		if maxf > res {
			res = maxf
		}
	}
	return res
}
