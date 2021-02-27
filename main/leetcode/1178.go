package main

func findNumOfValidWords(words []string, puzzles []string) []int {
	mp := make(map[int]int)
	for _, v := range words {
		bits := getBits(v)
		mp[bits]++
	}
	res := make([]int, len(puzzles))
	for k, v := range puzzles {
		puzzleBit := getBits(v)
		first := getBits(string(v[0]))

		n := puzzleBit
		for n > 0 {
			if (n&first != 0) && mp[n] > 0 {
				res[k] += mp[n]
			}
			n = (n - 1) & puzzleBit
		}
	}
	return res
}
func getBits(str string) int {
	res := 0
	for _, v := range str {
		offset := v - 'a'
		status := 1 << offset
		res = res | status
	}
	return res
}
