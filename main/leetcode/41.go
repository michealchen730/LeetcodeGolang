package main

func firstMissingPositive(nums []int) int {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			if mp[nums[i]] == 0 {
				temp := mp[nums[i]-1] + mp[nums[i]+1] + 1
				mp[nums[i]-mp[nums[i]-1]], mp[nums[i]+mp[nums[i]+1]], mp[nums[i]] = temp, temp, temp
			}
		}
	}
	return 1 + mp[1]
}
