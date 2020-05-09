package main

func kLengthApart(nums []int, k int) bool {
	if k == 0 {
		return true
	}
	i := -1
	for j, v := range nums {
		if v == 1 {
			if i != -1 {
				if j-i <= k {
					return false
				}
			}
			i = j
		}
	}
	return true
}
