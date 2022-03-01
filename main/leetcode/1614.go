package main

func maxDepth1614(s string) int {
	res := 0
	t := 0
	for _, v := range s {
		if v == '(' {
			t++
		}
		if v == ')' {
			if t > res {
				res = t
			}
			t--
		}
	}
	return res
}
