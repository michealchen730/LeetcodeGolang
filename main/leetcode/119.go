package main

func getRow(rowIndex int) []int {
	var res []int
	t := 0
	for t < rowIndex+1 {
		temp := 1

		for i := 1; i < len(res); i++ {
			next := res[i]
			res[i] += temp
			temp = next
		}
		res = append(res, 1)
		t++
	}
	return res
}
