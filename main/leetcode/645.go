package main

func findErrorNums(nums []int) []int {
	x := 0
	for _, v := range nums {
		x ^= v
	}
	for i := 1; i <= len(nums); i++ {
		x ^= i
	}
	lowbit := x & (-x)
	n1, n2 := 0, 0
	for _, v := range nums {
		if v&lowbit == 0 {
			n1 ^= v
		} else {
			n2 ^= v
		}
	}
	for i := 1; i <= len(nums); i++ {
		if i&lowbit == 0 {
			n1 ^= i
		} else {
			n2 ^= i
		}
	}
	for _, v := range nums {
		if v == n1 {
			return []int{n1, n2}
		}
	}
	return []int{n2, n1}
}
