package main

func longestDecomposition(text string) int {
	res := 0
	p1, p2 := 0, 0
	p3, p4 := len(text)-1, len(text)-1
	for p2 < p3 {
		if text[p1:p2+1] == text[p3:p4+1] {
			res += 2
			p1, p4 = p2+1, p3-1
		}
		p2++
		p3--
	}
	if p1 <= p4 {
		res++
	}
	return res
}
