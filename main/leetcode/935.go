package main

const m = 1000000007

func knightDialer(n int) int {
	if n == 1 {
		return 10
	}
	arr := make([]int, 9)
	arr[0] = 2
	arr[1] = 2
	arr[2] = 2
	arr[3] = 2
	arr[4] = 3
	arr[5] = 3
	arr[6] = 2
	arr[7] = 2
	arr[8] = 2
	for n > 2 {
		arr2 := make([]int, 9)
		arr2[0] = (arr[4] + arr[5]) % m
		arr2[1] = (arr[5] + arr[7]) % m
		arr2[2] = (arr[6] + arr[8]) % m
		arr2[3] = (arr[4] + arr[7]) % m
		arr2[4] = (arr[3] + arr[8] + arr[0]) % m
		arr2[5] = (arr[0] + arr[1] + arr[6]) % m
		arr2[6] = (arr[2] + arr[5]) % m
		arr2[7] = (arr[1] + arr[3]) % m
		arr2[8] = (arr[2] + arr[4]) % m
		arr = arr2
		n--
	}
	sum := 0
	for _, v := range arr {
		sum = (sum + v) % m
	}
	return sum

}
