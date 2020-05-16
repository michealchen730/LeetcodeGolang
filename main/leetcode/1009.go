package main

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}
	//先换十进制
	var arr []int
	for N != 0 {
		t := N % 2
		arr = append(arr, t^1)
		N = N / 2
	}
	res, temp := 0, 1
	for i := 0; i < len(arr); i++ {
		res += temp * arr[i]
		temp = temp << 1
	}
	return res
}
