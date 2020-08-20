package main

func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	fib1, fib2 := 0, 1
	for N > 1 {
		fib2, fib1 = fib2+fib1, fib2
		N--
	}
	return fib2
}
