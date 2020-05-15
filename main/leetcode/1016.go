package main

import (
	"strings"
)

func queryString(S string, N int) bool {
	for i := 1; i <= N; i++ {
		var bytes []byte
		num := i
		for num != 0 {
			t := num % 2
			num = num / 2
			bytes = append(bytes, byte(t+48))
		}
		reverseBytes1016(bytes)
		if !strings.Contains(S, string(bytes)) {
			return false
		}
	}
	return true
}

func reverseBytes1016(arr []byte) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
