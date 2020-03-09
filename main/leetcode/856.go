package main

import (
	"math"
)

func scoreOfParentheses(S string) int {
	//var arr []byte
	res, flag := 0, 0
	p := 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			p++
			//arr=append(arr,'(')
			flag = 1
		} else {
			//if len(arr)==len(S)-i{
			//	res+=int(math.Pow(2,float64(len(arr)-1)))
			//	break
			//}
			if p == len(S)-i {
				res += int(math.Pow(2, float64(p-1)))
				break
			}
			//arr=arr[1:]
			p--
			if flag == 1 {
				//res+=int(math.Pow(2,float64(len(arr))))
				res += int(math.Pow(2, float64(p)))
				flag = 0
			}
		}
	}
	return res
}
