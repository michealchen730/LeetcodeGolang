package main

func findBestValue(arr []int, target int) int {
	//先看这么一个假设，例子如下，target=15 arr=[1,2,3]，这个例子里，我们的极限也就只能是让数组保持不动
	//那么，根据这个例子，二分查找的上限是数组中最大的那个值
	low, high, sum := 0, 0, 0
	for _, v := range arr {
		high = max(high, v)
		sum += v
	}
	res, diff := high, abs(sum-target)
	if sum-target <= 0 {
		return res
	}
	for low <= high {
		mid := low + (high-low)/2
		temp := 0
		for _, v := range arr {
			if v > mid {
				temp += mid
			} else {
				temp += v
			}
		}
		if abs(temp-target) < diff {
			diff = abs(temp - target)
			res = mid
		} else if abs(temp-target) == diff {
			res = min(mid, res)
		}
		if temp-target == 0 {
			return mid
		} else if temp-target > 0 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return res
}
