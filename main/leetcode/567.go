package main

import "sort"

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	n := len(s1)
	sum := 0
	tmp := 0
	for i := 0; i < n; i++ {
		sum += int(s1[i] - 'a')
		tmp += int(s2[i] - 'a')
	}
	if tmp == sum && compareStrings(s1, s2[:n]) {
		return true
	}
	for j := n; j < len(s2); j++ {
		tmp += int(s2[j] - 'a')
		tmp -= int(s2[j-n] - 'a')
		if tmp == sum && compareStrings(s1, s2[j-n+1:j+1]) {
			return true
		}
	}
	return false
}

func compareStrings(s1, s2 string) bool {
	b1, b2 := []byte(s1), []byte(s2)
	sort.Slice(b1, func(i, j int) bool {
		return b1[i] < b1[j]
	})
	sort.Slice(b2, func(i, j int) bool {
		return b2[i] < b2[j]
	})
	if string(b1) == string(b2) {
		return true
	}
	return false
}
