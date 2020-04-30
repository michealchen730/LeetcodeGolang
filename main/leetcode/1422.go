package main

func maxScore(s string) int {
	res := 0
	nums0, nums1 := 0, 0
	for _, v := range s {
		if v == '0' {
			nums0++
		}
		if v == '1' {
			nums1++
		}
	}
	left, right := 0, nums1
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			left++
		} else {
			right--
		}
		if left+right > res {
			res = left + right
		}
	}
	return res
}
