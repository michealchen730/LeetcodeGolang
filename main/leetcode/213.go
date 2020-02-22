package main

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}

	res := make([]int, len(nums))
	res[0] = nums[0]
	if nums[0] > nums[1] {
		res[1] = nums[0]
	} else {
		res[1] = nums[1]
	}

	arr := make([]int, len(nums))
	arr[1] = nums[1]

	i := 2
	for i < len(res) {
		if nums[i]+res[i-2] > res[i-1] {
			res[i] = nums[i] + res[i-2]
		} else {
			res[i] = res[i-1]
		}
		if nums[i]+arr[i-2] > arr[i-1] {
			arr[i] = nums[i] + arr[i-2]
		} else {
			arr[i] = arr[i-1]
		}
		i++
	}
	if res[i-1] == res[i-2] { //最末尾的房子没动
		return res[i-1]
	} else {
		if res[i-1] == arr[i-1] { //最开始的房子没动
			return res[i-1]
		} else {
			if res[i-2] > arr[i-1] {
				return res[i-2]
			} else {
				return arr[i-1]
			}
		}
	}
}
