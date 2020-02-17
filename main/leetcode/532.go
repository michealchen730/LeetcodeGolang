package main

import (
	"sort"
)

func findPairs(nums []int, k int) int {
	tmap:=make(map[int]int)
	sort.Ints(nums)
	res:=0
	for i:=0;i<len(nums)-1;i++{
		BR:
		for j:=i+1;j<len(nums);j++{
			if nums[j]-nums[i]==k {
				if tmap[nums[j]]==0 {
					tmap[nums[j]]=1
					res++
				}
			}
			if nums[j]-nums[i]>k {
				break BR
			}
		}
	}
	return res
}
