package main

func maxNumberOfBalloons(text string) int {
	res := make([]int, 5)
	for i := 0; i < len(text); i++ {
		switch text[i] {
		case 'b':
			res[0] += 2
		case 'a':
			res[1] += 2
		case 'l':
			res[2] += 1
		case 'o':
			res[3] += 1
		case 'n':
			res[4] += 2
		}
	}
	temp := res[0]
	for i := 1; i < len(res); i++ {
		if res[i] < temp {
			temp = res[i]
		}
	}
	return temp / 2
}
