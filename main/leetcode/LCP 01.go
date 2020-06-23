package main

func game(guess []int, answer []int) int {
	res := 0
	for k, v := range guess {
		if v == answer[k] {
			res++
		}
	}
	return res
}
