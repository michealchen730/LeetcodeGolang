package main

//最简单的暴力做法
func twoSum01(nums []int, target int) []int {
	var res []int = []int{0, 0}
	length := len(nums)
EX:
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				res[0] = i
				res[1] = j
				break EX
			}
		}
	}
	return res
}

//利用哈希
func twoSum02(nums []int, target int) []int {
	mp := make(map[int]int)
	for k, v := range nums {
		if _, ok := mp[target-v]; ok {
			return []int{mp[target-v], k}
		}
		mp[v] = k
	}
	return []int{}
}
