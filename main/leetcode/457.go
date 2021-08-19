package main

func circularArrayLoop(nums []int) bool {
	//for循环完成一次遍历
	n := len(nums)
	for k, _ := range nums {
		//把所有超出数组长度的跳转次数都处理下
		nums[k] %= n
		//把所有一次循环全部排除
		if nums[k] == 0 {
			nums[k] = n
		}
	}
	//第二次遍历，start用来代表这次循环的唯一标识
	start := n
	for k, v := range nums {
		//v<n说明是合法的起点
		if v < n {
			//start++保持唯一性
			start++
			//循环全正还是全负的标识
			positive := v > 0
			nums[k] = start
			for {
				nextIndex := (k + v) % n
				if nextIndex < 0 {
					nextIndex += n
				}
				if nums[nextIndex] == start {
					return true
				} else {
					if nums[nextIndex] > 0 != positive || nums[nextIndex] >= n {
						break
					}
					k, v = nextIndex, nums[nextIndex]
				}
				nums[nextIndex] = start
			}
		}
	}
	return false
}
