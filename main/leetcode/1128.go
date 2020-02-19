package main

func numEquivDominoPairs(dominoes [][]int) int {
	res := make([]int, 100)
	max, min, temp := 0, 0, 0
	for i := 0; i < len(dominoes); i++ {
		if dominoes[i][0] < dominoes[i][1] {
			min, max = dominoes[i][0], dominoes[i][1]
		} else {
			min, max = dominoes[i][1], dominoes[i][0]
		}
		temp = min*10 + max
		res[temp]++
	}
	sum := 0
	for i := 11; i < len(res); i++ {
		if res[i] != 0 {
			sum += res[i] * (res[i] - 1) / 2
		}
	}
	return sum
}
