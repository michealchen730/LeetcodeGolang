package main

import (
	"strings"
)

//竞赛当天写的是个特暴力的方法
func sortString(s string) string {
	arr := make([]int, 26)
	for _, v := range s {
		arr[v-'a']++
	}
	var res strings.Builder
	t, turn := 0, 0
	for t != len(s) {
		if turn%2 == 0 {
			for k, v := range arr {
				if v != 0 {
					arr[k]--
					t++
					res.WriteByte(byte(k + 'a'))
				}
			}
		} else {
			for k := 25; k >= 0; k-- {
				if arr[k] != 0 {
					arr[k]--
					t++
					res.WriteByte(byte(k + 'a'))
				}
			}
		}
		turn++
	}
	return res.String()
}
