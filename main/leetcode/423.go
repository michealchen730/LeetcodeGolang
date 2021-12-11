package main

import (
	"strings"
)

//func originalDigits(s string) string {
//	arr:=make([]int,10)
//	for k, _ := range s {
//		switch s[k] {
//		case 'z':
//			arr[0]++
//		case 'w':
//			arr[2]++
//		case 'u':
//			arr[4]++
//		case 'x':
//			arr[6]++
//		case 'g':
//			arr[8]++
//		case 'f':
//			arr[5]++
//		case 'v':
//			arr[7]++
//		case 'h':
//			arr[3]++
//		case 'o':
//			arr[1]++
//		case 'i':
//			arr[9]++
//		default:
//		}
//	}
//	arr[5]-=arr[4]
//	arr[7]-=arr[5]
//	arr[3]-=arr[8]
//	arr[1]-=arr[0]+arr[2]+arr[4]
//	arr[9]-=arr[6]+arr[8]+arr[5]
//	var res strings.Builder
//	for k, v := range arr {
//		for i:=0;i<v;i++{
//			res.WriteByte(byte(48 + k))
//		}
//	}
//	return res.String()
//}

func originalDigits(s string) string {
	arr := make([]int, 11)
	mp := map[byte]int{'z': 0, 'w': 2, 'u': 4, 'x': 6, 'g': 8, 'f': 5, 'v': 7, 'h': 3, 'o': 1, 'i': 9, 'e': 10, 'n': 10, 's': 10, 'r': 10, 't': 10}
	for k, _ := range s {
		arr[mp[s[k]]]++
	}
	arr[5] -= arr[4]
	arr[7] -= arr[5]
	arr[3] -= arr[8]
	arr[1] -= arr[0] + arr[2] + arr[4]
	arr[9] -= arr[6] + arr[8] + arr[5]
	var res strings.Builder
	for k, v := range arr[:10] {
		for i := 0; i < v; i++ {
			res.WriteByte(byte(48 + k))
		}
	}
	return res.String()
}
