package main

func addDigits(num int) int {
	for num >= 10 {
		t := 0
		for num != 0 {
			t += num % 10
			num /= 10
		}
		num = t
	}
	return num
}
