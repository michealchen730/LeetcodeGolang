package main

func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) == len(str2) {
		if str1 == str2 {
			return str1
		} else {
			return ""
		}
	}
	if len(str1) > len(str2) {
		if str1[0:len(str2)] == str2 {
			return gcdOfStrings(str1[len(str2):], str2)
		} else {
			return ""
		}
	}
	if len(str2) > len(str1) {
		if str2[0:len(str1)] == str1 {
			return gcdOfStrings(str2[len(str1):], str1)
		} else {
			return ""
		}
	}
	return ""
}
