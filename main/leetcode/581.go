package main

import (
	"math"
)

//func findUnsortedSubarray(nums []int) int {
//	nums2:=make([]int,len(nums))
//	copy(nums2,nums)
//	sort.Ints(nums2)
//	start:=0
//	end:=len(nums2)-1
//	for start<=end{
//		if nums2[start]==nums[start]{
//			start++
//		}else{
//			break
//		}
//	}
//	if start==end+1{
//		return 0
//	}
//	for end>start{
//		if nums2[end]==nums[end]{
//			end--
//		}else{
//			break
//		}
//	}
//	return end-start+1
//}

func findUnsortedSubarray(nums []int) int {
	minn := math.MaxInt32
	maxn := math.MinInt32
	left := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < minn {
			minn = nums[i]
		} else {
			left = i
		}
	}
	if left == len(nums)-1 {
		return 0
	}
	right := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxn {
			maxn = nums[i]
		} else {
			right = i
		}
	}
	return right - left + 1
}
