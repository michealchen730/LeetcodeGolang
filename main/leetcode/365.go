package main

func canMeasureWater(x int, y int, z int) bool {
	if z == 0 {
		return true
	}

	if z > x+y {
		return false
	}
	temp := getGreatestCommonDivisor(x, y)
	if z%temp == 0 {
		return true
	}
	return false
}

func getGreatestCommonDivisor(x, y int) int {
	if x > y {
		x, y = y, x
	}
	var rem int
	for y > 0 {
		rem = x % y
		x = y
		y = rem
	}
	return x
}
