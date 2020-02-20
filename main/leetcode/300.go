package main

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	arr := make([]int, len(nums))
	arr[0] = 1
	res := 1
	for i := 1; i < len(arr); i++ {
		arr[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && arr[j]+1 > arr[i] {
				arr[i] = arr[j] + 1
			}
		}
		if arr[i] > res {
			res = arr[i]
		}
	}
	return res
}
