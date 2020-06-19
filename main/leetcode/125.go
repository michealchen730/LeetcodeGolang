package main

func isPalindrome125(s string) bool {
	//只是图效率没太大意义
	var bytes []byte
	for _, v := range s {
		if (v >= 48 && v <= 57) || (v >= 97 && v <= 122) {
			bytes = append(bytes, byte(v))
		}
		if v >= 65 && v <= 90 {
			bytes = append(bytes, byte(v+32))
		}
	}
	i, j := 0, len(bytes)-1
	for i < j {
		if bytes[i] != bytes[j] {
			return false
		}
		i++
		j--
	}
	return true
}
