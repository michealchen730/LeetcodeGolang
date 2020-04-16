package main

func numSteps(s string) int {
	res := 0
	for len(s) != 1 {
		if s[len(s)-1] == '1' {
			bytes := []byte{'0'}
			leap := 1
			for i := len(s) - 2; i >= 0; i-- {
				temp := int(s[i]-'0') + leap
				if temp == 1 {
					bytes = append(bytes, '1')
					leap = 0
				} else if temp == 2 {
					bytes = append(bytes, '0')
					leap = 1
				} else {
					bytes = append(bytes, '0')
					leap = 0
				}
			}
			if leap == 1 {
				bytes = append(bytes, '1')
			}
			h, t := 0, len(bytes)-1
			for h < t {
				bytes[h], bytes[t] = bytes[t], bytes[h]
				h++
				t--
			}
			s = string(bytes)
		} else {
			s = s[:len(s)-1]
		}
		res++
	}
	return res
}
