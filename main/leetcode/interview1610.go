package main

func maxAliveYear(birth []int, death []int) int {
	arr := make([]int, 103)
	live := 0
	res := 1900
	for k, v := range birth {
		updateRange(v-1900+1, death[k]-1900+1, 1, arr)
	}

	for i := 0; i < 103; i++ {
		if live < find(i, arr) {
			live = find(i, arr)
			res = 1900 + i - 1
		}
	}
	return res
}
