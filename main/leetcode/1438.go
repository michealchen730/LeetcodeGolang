package main

func longestSubarray(nums []int, limit int) int {
	res := 1
	for i := 0; i < len(nums); i++ {
		if len(nums)-i <= res {
			break
		}
		mx := nums[i]
		mn := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > mx {
				mx = nums[j]
			}
			if nums[j] < mn {
				mn = nums[j]
			}
			if abs(mx-mn) <= limit {
				if j-i+1 > res {
					res = j - i + 1
				}
			} else {
				break
			}
		}
	}
	return res
}
