package main

func generateTheString(n int) string {
	s := ""
	if n%2 == 0 {
		for i := 1; i < n; i++ {
			s += "a"
		}
		s += "b"
	} else {
		for i := 0; i < n; i++ {
			s += "a"
		}
	}
	return s
}
