package main

func addToArrayForm(A []int, K int) []int {
	B := transToArr(K)
	var res []int
	i, j := len(A)-1, len(B)-1
	flag := 0
	for i >= 0 && j >= 0 {
		v := A[i] + B[j] + flag
		if v >= 10 {
			v %= 10
			flag = 1
		} else {
			flag = 0
		}
		i--
		j--
		res = append([]int{v}, res...)
	}

	for i != -1 {
		v := A[i] + flag
		if v >= 10 {
			v %= 10
			flag = 1
		} else {
			flag = 0
		}
		i--
		res = append([]int{v}, res...)
	}
	for j != -1 {
		v := B[j] + flag
		if v >= 10 {
			v %= 10
			flag = 1
		} else {
			flag = 0
		}
		j--
		res = append([]int{v}, res...)
	}
	if flag != 0 {
		res = append([]int{flag}, res...)
	}

	return res
}
func transToArr(n int) []int {
	var res []int
	for n != 0 {
		t := n % 10
		res = append([]int{t}, res...)
		n /= 10
	}
	return res
}
