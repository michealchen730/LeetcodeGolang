package main

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	arr := make([]int, len(nums))
	arr[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+nums[i]; j++ {
			if arr[j] == 0 || min(arr[j], arr[i]+1) == arr[i]+1 {
				arr[j] = arr[i] + 1
			}
		}
	}
	return arr[len(arr)-1]
}
