package main

func getNoZeroIntegers(n int) []int {
	for i := 1; i < n; i++ {
		if haveZero(i) && haveZero(n-i) {
			return []int{i, n - i}
		}
	}
	return []int{0, 0}
}

func haveZero(n int) bool {
	for n != 0 {
		if n%10 == 0 {
			return false
		} else {
			n = n / 10
		}
	}
	return true
}
