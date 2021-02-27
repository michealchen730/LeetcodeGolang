package main

//逐步遍历
//没有利用到数组升序这一条件
//func search(nums []int, target int) int {
//	res:=0
//	for _,v:=range nums{
//		if v==target{
//			res++
//		}
//	}
//	return res
//}
func searchO53(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	//如果存在该值，返回最左边的位置
	//如果不存在该值，返回比这个值大的值的第一个位置
	binarySearch := func(tar int) int {
		if tar > nums[len(nums)-1] {
			return len(nums)
		}
		low, high := 0, len(nums)-1
		for low < high {
			mid := low + (high-low)/2
			if nums[mid] < tar {
				low = mid + 1
			} else if nums[mid] > tar {
				high = mid - 1
			} else {
				high = mid
			}
		}
		return low
	}
	left := binarySearch(target)
	right := binarySearch(target + 1)
	if left == len(nums) || nums[left] != target {
		return 0
	}
	return right - left
}
