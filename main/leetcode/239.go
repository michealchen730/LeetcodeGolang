package main

func maxSlidingWindow(nums []int, k int) []int {
	res, t := 0, k
	for i := 0; i < k; i++ {
		res += nums[i]
	}
	ans := res
	for t < len(nums) {
		res += nums[t]
		res -= nums[t-k]
		t++
		ans = max(ans, res)
	}
	return ans
}
