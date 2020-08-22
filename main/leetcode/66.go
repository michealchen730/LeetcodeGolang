package main

func plusOne(digits []int) []int {
	carry := 0
	for i := len(digits) - 1; i >= 0; i-- {
		sum := carry + digits[i]
		if i == len(digits)-1 {
			sum += 1
		}
		carry = sum / 10
		sum %= 10
		digits[i] = sum
	}
	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}
