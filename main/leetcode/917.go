package main

import "fmt"

func reverseOnlyLetters(S string) string {
	bytes := []byte(S)
	pt1, pt2 := 0, len(S)-1
	for pt1 < pt2 {
		if ((bytes[pt1] >= 'A' && bytes[pt1] <= 'Z') || (bytes[pt1] >= 'a' && bytes[pt1] <= 'z')) && ((bytes[pt2] >= 'A' && bytes[pt2] <= 'Z') || (bytes[pt2] >= 'a' && bytes[pt2] <= 'z')) {
			bytes[pt1], bytes[pt2] = bytes[pt2], bytes[pt1]
			pt1++
			pt2--
		}
		if !((bytes[pt1] >= 'A' && bytes[pt1] <= 'Z') || (bytes[pt1] >= 'a' && bytes[pt1] <= 'z')) {
			pt1++
		}
		if !((bytes[pt2] >= 'A' && bytes[pt2] <= 'Z') || (bytes[pt2] >= 'a' && bytes[pt2] <= 'z')) {
			pt2--
		}
	}
	return string(bytes)
}

func main() {
	fmt.Println(reverseOnlyLetters("Test1ng-Leet=code-Q!"))
}
