package main

func isPerfectSquare(num int) bool {
	left, right := 0, num
	for left <= right {
		mid := (left + right) / 2
		s := mid * mid
		if s < num {
			left = mid + 1
		} else if s > num {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}
