package main

func isNumber65(s string) bool {
	state := 0
	matrix := make([][]int, 10)
	matrix[0] = []int{0, 1, 2, -1, 4}
	matrix[1] = []int{-1, -1, 2, -1, 4}
	matrix[2] = []int{8, -1, 2, 3, 9}
	matrix[3] = []int{-1, 5, 6, -1, -1}
	matrix[4] = []int{-1, -1, 7, -1, -1}
	matrix[5] = []int{-1, -1, 6, -1, -1}
	matrix[6] = []int{8, -1, 6, -1, -1}
	matrix[7] = []int{8, -1, 7, 3, -1}
	matrix[8] = []int{8, -1, -1, -1, -1}
	matrix[9] = []int{8, -1, 7, 3, -1}
	mapper := func(input byte) int {
		if input == ' ' {
			return 0
		}
		if input >= '0' && input <= '9' {
			return 2
		}
		if input == '+' || input == '-' {
			return 1
		}
		if input == 'e' || input == 'E' {
			return 3
		}
		if input == '.' {
			return 4
		}
		if input == 'E' {
			return 5
		}
		return -1
	}
	for k, _ := range s {
		if state == -1 || mapper(s[k]) == -1 {
			return false
		} else {
			state = matrix[state][mapper(s[k])]
		}
	}
	if state == 2 || state > 5 {
		return true
	}
	return false
}
