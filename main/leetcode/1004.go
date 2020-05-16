package main

func longestOnes(A []int, K int) int {
	i, j := 0, 0
	width := 0
	left := K - 1 + A[j]
	for j+1 < len(A) {
		j++
		if A[j] == 0 {
			left--
		}
		if left >= 0 {
			width = max(width, j-i+1)
		}
		for left < 0 {
			if A[i] == 0 {
				left++
			}
			i++
		}
	}
	return width
}
