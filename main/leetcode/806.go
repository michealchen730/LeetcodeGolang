package main

func numberOfLines(widths []int, s string) []int {
	res := 0
	temp := 0
	for _, v := range s {
		if widths[int(v)-int('a')]+temp > 100 {
			res++
			temp = widths[int(v)-int('a')]
		} else {
			temp += widths[int(v)-int('a')]
		}
	}
	if temp > 0 {
		res++
	}
	return []int{res, temp}
}
