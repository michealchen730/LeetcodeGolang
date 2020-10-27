package main

func findMaxLength(nums []int) int {
	mp := map[int]int{0: -1}
	cnt0 := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			cnt0++
		}
		if v, ok := mp[i+1-cnt0-cnt0]; ok {
			res = max(res, i-v)
		} else {
			mp[i+1-cnt0-cnt0] = i
		}
	}
	return res
}
