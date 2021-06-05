package main

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	for n != 1 {
		if n&1 == 1 {
			return false
		} else {
			n >>= 1
		}
	}
	return true
}
