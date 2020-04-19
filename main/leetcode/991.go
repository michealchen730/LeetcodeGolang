package main

func brokenCalc(X int, Y int) int {
	if X == Y {
		return 0
	} else if X > Y {
		return X - Y
	} else {
		temp := 1
		if Y%2 == 0 {
			Y = Y / 2
		} else {
			Y = Y/2 + 1
			temp++
		}
		return temp + brokenCalc(X, Y)
	}
}
