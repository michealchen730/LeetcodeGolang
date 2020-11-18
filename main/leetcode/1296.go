package main

import (
	"sort"
)

func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}

	mp := make(map[int]int)

	sort.Ints(nums)
	var arr []int

	for i, v := range nums {
		if i == 0 || nums[i] != nums[i-1] {
			arr = append(arr, nums[i])
		}
		mp[v]++
	}

	for i := 0; i < len(arr); i++ {
		if mp[arr[i]] == 0 {
			continue
		} else if mp[arr[i]] < 0 {
			return false
		} else {
			cnt := mp[arr[i]]
			for n := arr[i]; n < arr[i]+k; n++ {
				mp[n] -= cnt
				if mp[n] < 0 {
					return false
				}
			}
		}
	}
	return true

}
