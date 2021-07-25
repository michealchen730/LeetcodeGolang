package main

func isCovered(ranges [][]int, left int, right int) bool {
	diff := make([]int, 52)
	for _, v := range ranges {
		diff[v[0]]++
		diff[v[1]+1]--
	}
	sum := 0
	for i := 0; i <= right; i++ {
		sum += diff[i]
		if sum <= 0 && i >= left {
			return false
		}
	}
	return true
}
