package main

func numSubarraysWithSum(nums []int, goal int) int {
	mp := map[int]int{0: 1}
	tmp := 0
	res := 0
	for _, v := range nums {
		tmp += v
		res += mp[tmp-goal]
		mp[tmp-goal]++
	}
	return res
}
