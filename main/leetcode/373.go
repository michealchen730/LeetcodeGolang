package main

import "math"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	//k有可能会超出两个数组长度的乘积，因此需要调整k的大小
	m, n := len(nums1), len(nums2)
	if k > m*n {
		k = m * n
	}
	var res [][]int
	flag := true
	if m > n {
		m, n = n, m
		nums1, nums2 = nums2, nums1
		flag = false
	}
	points := make([]int, m)
	for len(res) < k {
		mn := math.MaxInt32
		idx := 0
		for i, _ := range nums1 {
			if points[i] >= n {
				continue
			}
			if nums2[points[i]]+nums1[i] < mn {
				mn = nums2[points[i]] + nums1[i]
				idx = i
			}
		}
		var element []int
		if flag {
			element = append(element, []int{nums1[idx], nums2[points[idx]]}...)
		} else {
			element = append(element, []int{nums2[points[idx]], nums1[idx]}...)
		}
		res = append(res, element)
		points[idx]++
	}
	return res
}
