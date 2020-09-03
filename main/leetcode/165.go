package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	s1 := strings.Split(version1, ".")
	s2 := strings.Split(version2, ".")

	k1, k2 := 0, 0

	for k1 < len(s1) || k2 < len(s2) {
		num1, num2 := 0, 0
		if k1 < len(s1) {
			num1, _ = strconv.Atoi(s1[k1])
		}
		if k2 < len(s2) {
			num2, _ = strconv.Atoi(s2[k2])
		}
		if num1 < num2 {
			return -1
		} else if num1 > num2 {
			return 1
		}
		k1++
		k2++
	}
	return 0
}
