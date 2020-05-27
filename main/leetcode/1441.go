package main

func buildArray(target []int, n int) []string {
	var res []string
	num := 0
	pos := 0
	for pos < len(target) {
		res = append(res, "Push")
		num++
		if num == target[pos] {
			pos++
		} else {
			res = append(res, "Pop")
		}
	}
	return res
}
