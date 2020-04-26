package main

import "strings"

func reformat(s string) string {
	if len(s) <= 1 {
		return s
	}
	var alp []int32
	var num []int32
	for _, v := range s {
		if v >= '0' && v <= '9' {
			num = append(num, v)
		} else {
			alp = append(alp, v)
		}
	}
	if abs(len(alp)-len(num)) <= 1 {
		var res strings.Builder
		if len(alp) > len(num) {
			for i := 0; i < len(num); i++ {
				res.WriteByte(byte(alp[i]))
				res.WriteByte(byte(num[i]))
			}
			res.WriteByte(byte(alp[len(alp)-1]))
		} else if len(alp) < len(num) {
			for i := 0; i < len(alp); i++ {
				res.WriteByte(byte(num[i]))
				res.WriteByte(byte(alp[i]))
			}
			res.WriteByte(byte(num[len(num)-1]))
		} else {
			for i := 0; i < len(alp); i++ {
				res.WriteByte(byte(num[i]))
				res.WriteByte(byte(alp[i]))
			}
		}
		return res.String()
	} else {
		return ""
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
