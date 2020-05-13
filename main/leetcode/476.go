package main

func findComplement(num int) int {
	//先换十进制
	var arr []int
	for num != 0 {
		t := num % 2
		arr = append(arr, t^1)
		num = num / 2
	}
	res, temp := 0, 1
	for i := 0; i < len(arr); i++ {
		res += temp * arr[i]
		temp = temp << 1
	}
	return res
}
