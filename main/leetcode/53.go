package main

func maxSubArray(nums []int) int {
	flag := 0
	maxs := nums[0]
	for i := 0; i < len(nums); i++ {
		maxs = max(nums[i], maxs)
		if nums[i] >= 0 {
			flag = 1
			break
		}
	}
	if flag == 0 {
		return maxs
	} else {
		maxs = 0
		temps, i := 0, 0
		for i < len(nums) {
			if temps+nums[i] >= 0 {
				temps += nums[i]
				if temps > maxs {
					maxs = temps
				}
			} else {
				temps = 0
			}
			i++
		}
		return maxs
	}
}
