package main

func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if i > k {
			delete(mp, nums[i-k-1])
		}
		if mp[nums[i]] {
			return true
		}
		mp[nums[i]] = true
	}
	return false
}
