package main

func minOperations(n int) int {
	if n%2 != 0 {
		return (n/2)*2 + (n/2)*(n/2-1)
	} else {
		return (n/2)*1 + (n/2)*(n/2-1)
	}
}
