package main

func nextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) == 1 {
		return []int{-1}
	}
	var arr []int
	res := make([]int, len(nums))
	for t := (len(nums))*2 - 1; t >= 0; t-- {
		i := t % len(nums)
		if len(arr) == 0 {
			arr = append(arr, nums[i])
			res[i] = -1
		} else {
			if nums[i] < arr[len(arr)-1] {
				res[i] = arr[len(arr)-1]
				arr = append(arr, nums[i])
			} else {
				for len(arr) != 0 {
					if arr[len(arr)-1] <= nums[i] {
						arr = arr[:len(arr)-1]
					} else {
						res[i] = arr[len(arr)-1]
						arr = append(arr, nums[i])
						break
					}
				}
				if len(arr) == 0 {
					arr = append(arr, nums[i])
					res[i] = -1
				}
			}
		}
	}
	return res
}
