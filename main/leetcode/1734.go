package main

func decode(encoded []int) []int {
	n := len(encoded) + 1
	total := 0
	for i := 1; i <= n; i++ {
		total ^= i
		if i < n-1 && i%2 != 0 {
			total ^= encoded[i]
		}
	}
	res := make([]int, n)
	res[0] = total
	for i := 1; i < n; i++ {
		res[i] = res[i-1] ^ encoded[i-1]
	}
	return res
}
