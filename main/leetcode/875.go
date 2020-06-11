package main

func minEatingSpeed(piles []int, H int) int {
	left, right := 1, 0
	for _, v := range piles {
		right = max(right, v)
	}
	ans := right
	for left <= right {
		mid, cnt := left+(right-left)/2, 0
		for _, v := range piles {
			cnt += (v-1)/mid + 1
		}
		if cnt <= H {
			ans = min(mid, ans)
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
