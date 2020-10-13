package main

//先试动态规划
func maxSlidingWindow(nums []int, k int) []int {
	left, right := make([]int, len(nums)), make([]int, len(nums))
	pieces := len(nums)/k - 1
	if len(nums)%k != 0 {
		pieces++
	}
	for i := 0; i <= pieces; i++ {
		temp := nums[k*i]
		for j := k * i; j < min(k*(i+1), len(nums)); j++ {
			temp = max(temp, nums[j])
			left[j] = temp
		}
		end := min(k*(i+1)-1, len(nums)-1)
		temp2 := nums[end]
		for j := end; j >= k*i; j-- {
			temp2 = max(temp2, nums[j])
			right[j] = temp2
		}
	}
	res := make([]int, len(nums)-k+1)
	for i := 0; i < len(res); i++ {
		res[i] = max(right[i], left[i+k-1])
	}
	return res
}

//双端队列
//时间复杂度这块儿想不太通
func maxSlidingWindow2(nums []int, k int) []int {
	var deque []int
	res := make([]int, len(nums)-k+1)
	for i := 0; i < k; i++ {
		if len(deque) == 0 {
			deque = append(deque, i)
		} else {
			for j := len(deque) - 1; j >= 0; j-- {
				if nums[deque[j]] <= nums[i] {
					deque = deque[:len(deque)-1]
				} else {
					deque = append(deque, i)
					break
				}
			}
			if len(deque) == 0 {
				deque = append(deque, i)
			}
		}
	}
	res[0] = nums[deque[0]]

	for i := 1; i < len(nums)-k+1; i++ {
		//检查堆顶元素是否还有效
		if deque[0] < i {
			deque = deque[1:]
		}
		//进堆
		for len(deque) != 0 {
			if nums[deque[len(deque)-1]] <= nums[i+k-1] {
				deque = deque[:len(deque)-1]
			} else {
				deque = append(deque, i+k-1)
				break
			}
		}
		if len(deque) == 0 {
			deque = append(deque, i+k-1)
		}
		res[i] = nums[deque[0]]
	}
	return res
}
