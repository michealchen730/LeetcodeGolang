package main

func diStringMatch(aS string) []int {
	mn, mx := 0, len(aS)
	var res []int
	for _, v := range aS {
		if v == 'I' {
			res = append(res, mn)
			mn++
		}
		if v == 'D' {
			res = append(res, mx)
			mx--
		}
	}
	res = append(res, mx)
	return res
}
