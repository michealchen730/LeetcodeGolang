package main

func printNumbers(n int) []int {
	mx := 1
	for i := 0; i < n; i++ {
		mx *= 10
	}
	res := make([]int, mx-1)
	for i := 0; i < len(res); i++ {
		res[i] = i + 1
	}
	return res
}
