package main

func shipWithinDays(weights []int, D int) int {
	left, right := 0, 0
	for _, v := range weights {
		right += v
		left = max(left, v)
	}
	ans := right
	for left <= right {
		mid := left + (right-left)/2
		sum, cnt := 0, 1
		for _, v := range weights {
			if sum+v > mid {
				cnt++
				sum = v
			} else {
				sum += v
			}
		}
		if cnt <= D {
			ans = min(mid, ans)
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
