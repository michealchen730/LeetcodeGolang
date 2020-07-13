package main

func multiplyI0805(A int, B int) int {
	if A < B {
		A, B = B, A
	}
	for B > 0 {
		return A + multiplyI0805(A, B-1)
	}
	return 0
}
