package main

func reorderedPowerOf2(N int) bool {
	arr := spiltN(N)
	temp := 1
	for temp <= 1000000000 {
		arr2 := spiltN(temp)
		flag := true
		for k, v := range arr {
			if v != arr2[k] {
				flag = false
				break
			}
		}
		if flag {
			return flag
		}
		temp *= 2
	}
	return false
}
func spiltN(N int) []int {
	arr := make([]int, 10)
	for N != 0 {
		t := N % 10
		arr[t]++
		N = N / 10
	}
	return arr
}
