package main

func numTimesAllBlue(light []int) int {
	temp := 0
	res := 0
	for i := 0; i < len(light); i++ {
		if light[i] > temp {
			temp = light[i]
		}
		if temp == i+1 {
			res++
		}
	}
	return res
}
