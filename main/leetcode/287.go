package main

func findDuplicate(nums []int) int {
	slow := 0
	fast := 0
	for true {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	res := 0
	for true {
		res = nums[res]
		slow = nums[slow]
		if slow == res {
			return res
		}
	}
	return res
}

func findDuplicate2(nums []int) int {
	temp := len(nums) - 1
	low, high := 1, temp
	for low < high {
		mid := low + (high-low)/2
		amount := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] >= low && nums[i] <= mid {
				amount++
			}
		}
		if amount > mid-low+1 {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
