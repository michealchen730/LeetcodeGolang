package main

func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	if n == 1 {
		return x
	}
	if n == 0 {
		return 1
	}
	temp1 := myPow(x, n/2)
	temp2 := 1.0
	if n%2 == 1 {
		temp2 = x
	}
	return temp1 * temp1 * temp2
}
