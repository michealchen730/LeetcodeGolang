package main

func matrixScore(A [][]int) int {
	for i := 0; i < len(A); i++ {
		if A[i][0] == 0 {
			reverseArr(A[i])
		}
	}
	for i := 0; i < len(A[0]); i++ {
		nums := 0
		for j := 0; j < len(A); j++ {
			if A[j][i] == 0 {
				nums++
			}
		}
		if nums > len(A)/2 {
			reverseMar(A, i)
		}
	}
	res := 0
	for i := 0; i < len(A); i++ {
		res += binToTen(A[i])
	}
	return res
}

func reverseArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			arr[i] = 0
		} else {
			arr[i] = 1
		}
	}
}
func reverseMar(mar [][]int, i int) {
	for j := 0; j < len(mar); j++ {
		if mar[j][i] == 0 {
			mar[j][i] = 1
		} else {
			mar[j][i] = 0
		}
	}
}
func binToTen(arr []int) int {
	res := 0
	temp := 1
	for i := len(arr) - 1; i >= 0; i-- {
		res += arr[i] * temp
		temp *= 2
	}
	return res
}
