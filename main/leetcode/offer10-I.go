package main

func fibO10(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	fib1, fib2 := 0, 1
	for N > 1 {
		fib2, fib1 = (fib2+fib1)%1000000007, fib2%1000000007
		N--
	}
	return fib2
}
