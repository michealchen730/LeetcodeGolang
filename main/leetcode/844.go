package main

func backspaceCompare(S string, T string) bool {
	if getStrings(S) == getStrings(T) {
		return true
	} else {
		return false
	}
}

func getStrings(S string) string {
	var res []byte
	for i := 0; i < len(S); i++ {
		if S[i] != '#' {
			res = append(res, S[i])
		} else {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		}
	}
	return string(res)
}
