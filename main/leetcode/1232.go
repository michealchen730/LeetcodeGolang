package main

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) == 2 {
		return true
	}
	x1, x2 := coordinates[1][0]-coordinates[0][0], coordinates[1][1]-coordinates[0][1]
	for k := 2; k < len(coordinates); k++ {
		if (coordinates[k][0]-coordinates[1][0])*x2 != x1*(coordinates[k][1]-coordinates[1][1]) {
			return false
		}
	}
	return true
}
