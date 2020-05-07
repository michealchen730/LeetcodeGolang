package main

func subtractProductAndSum(n int) int {
	product, sum := 1, 0
	for n != 0 {
		t := n % 10
		product *= t
		sum += t
		n /= 10
	}
	return product - sum
}
