package main

//暴力写法
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if j-i > k {
				break
			}
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}
	return false
}
