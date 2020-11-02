package main

func intersection(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int)
	for _, v := range nums1 {
		if mp[v] == 0 {
			mp[v]++
		}
	}
	var res []int
	for _, v := range nums2 {
		if mp[v] == 1 {
			mp[v] = 0
			res = append(res, v)
		}
	}
	return res
}
