package main

func subarraySum(nums []int, k int) int {
	mp := make(map[int]int)
	sum, res := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum == k {
			res++
		}
		if _, ok := mp[sum-k]; ok {
			res += mp[sum-k]
		}
		mp[sum]++
	}
	return res
}
