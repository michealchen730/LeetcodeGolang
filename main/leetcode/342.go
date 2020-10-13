package main

func isPowerOfFour(num int) bool {
	if num == 0 {
		return false
	}
	if num == 1 {
		return true
	}
	if num%4 != 0 {
		return false
	}
	return isPowerOfFour(num / 4)
}

func isPowerOfFour2(num int) bool {
	return num > 0 && num&(num-1) == 0 && num%3 == 1
}
