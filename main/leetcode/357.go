package main

func main() {
	countNumbersWithUniqueDigits(5)
}

func countNumbersWithUniqueDigits(n int) int {
	if n < 0 {
		return 10
	}
	arr1 := make([]int, 11)
	arr1[0], arr1[1] = 1, 10
	temp := 9
	flag := 9
	for i := 2; i < len(arr1); i++ {
		temp *= flag
		flag--
		arr1[i] = temp + arr1[i-1]
	}
	if n >= 10 {
		return arr1[10]
	} else {
		return arr1[n]
	}
}

func countNumbersWithUniqueDigits2(n int) int {
	if n < 0 {
		return 10
	}
	arr := []int{1, 10, 91, 739, 5275, 32491, 168571, 712891, 2345851, 5611771, 8877691}
	if n >= 10 {
		return arr[10]
	} else {
		return arr[n]
	}
}
