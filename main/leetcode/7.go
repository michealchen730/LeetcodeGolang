package main

import "strconv"

func reverse(x int) int {
	s := strconv.Itoa(x)
	sBytes := []byte(s)
	var i int
	j := len(sBytes) - 1
	if sBytes[0] == '-' {
		i = 1
	}
	for i < j {
		sBytes[i], sBytes[j] = sBytes[j], sBytes[i]
		i++
		j--
	}
	s = string(sBytes)
	temp, _ := strconv.ParseInt(s, 10, 64)
	if temp > 2147483647 || temp < (-2147483648) {
		temp = 0
	}
	s = strconv.FormatInt(temp, 10)
	res, _ := strconv.Atoi(s)
	return res
}
