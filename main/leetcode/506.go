package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	arr := make([]int, len(nums))
	copy(arr, nums)
	sort.Ints(arr)
	mp := make(map[int]int)
	for i := len(arr) - 1; i >= 0; i-- {
		mp[arr[i]] = len(arr) - i
	}
	res := make([]string, len(nums))
	for k, v := range nums {
		if mp[v] == 1 {
			res[k] = "Gold Medal"
			continue
		}
		if mp[v] == 2 {
			res[k] = "Silver Medal"
			continue
		}
		if mp[v] == 3 {
			res[k] = "Bronze Medal"
			continue
		}
		res[k] = strconv.Itoa(mp[v])
	}
	return res

}
