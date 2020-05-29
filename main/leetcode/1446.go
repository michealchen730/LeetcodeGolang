package main

func maxPower(s string) int {
	mx := 1
	temp := 1
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			temp = 1
		} else {
			temp++
			if temp > mx {
				mx = temp
			}
		}
	}
	return mx
}
