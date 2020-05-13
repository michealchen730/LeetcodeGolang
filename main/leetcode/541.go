package main

func reverseStr(s string, k int) string {
	if len(s) <= 1 {
		return s
	}
	bytes := []byte(s)
	reverseBytes(bytes, 0, k)
	return string(bytes)
}

func reverseBytes(arr []byte, i int, k int) {
	j := k - 1 + i
	//剩余字符串小于k
	if j > len(arr)-1 {
		j = len(arr) - 1
		swapBytes(arr, i, j)
		return
	}
	//剩余字符串大于k小于等于2k
	if j+k >= len(arr)-1 {
		swapBytes(arr, i, j)
		return
	} else {
		t := i + 2*k
		swapBytes(arr, i, j)
		reverseBytes(arr, t, k)
	}
}
func swapBytes(arr []byte, i int, j int) {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
