package main

func binaryGap(N int) int {
	var ints []int
	for N != 0 {
		ints = append(ints, N%2)
		N = N / 2
	}
	flag := -1
	res := 0
	for i := 0; i < len(ints); i++ {
		if ints[i] == 1 {
			if flag != -1 && res < i-flag {
				res = i - flag
			}
			flag = i
		}
	}
	return res
}
