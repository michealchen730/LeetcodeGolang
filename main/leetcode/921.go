package main

func minAddToMakeValid(S string) int {
	res, state := 0, 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			if state < 0 {
				res += -state
				state = 0
			}
			state++
		}
		if S[i] == ')' {
			state--
		}
	}
	if state > 0 {
		res += state
	} else {
		res += -state
	}
	return res
}
