package main

func minTimeLCP(time []int, m int) int {
	if m >= len(time) {
		return 0
	}
	left, right := time[0], 0
	for _, v := range time {
		right += v
		left = min(left, v)
	}
	ans := right
	for left <= right {
		mid := left + (right-left)/2
		sum, cnt, mx := 0, 1, 0
		for _, v := range time {
			mx = max(mx, v)
			if sum+v-mx > mid {
				cnt++
				sum, mx = v, v
			} else {
				sum += v
			}
		}
		if cnt <= m {
			ans = min(ans, mid)
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
