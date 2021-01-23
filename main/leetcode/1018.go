package main

func prefixesDivBy5(A []int) []bool {
	res := make([]bool, len(A))
	temp := 0
	for k, v := range A {
		temp += v
		if temp%5 == 0 {
			res[k] = true
		} else {
			res[k] = false
		}
		temp %= 10
		temp <<= 1

	}
	return res
}
