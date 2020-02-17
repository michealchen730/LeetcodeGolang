package main

import (
	"strconv"
)
func compress(chars []byte) int {
	if len(chars)<=1 {
		return len(chars)
	}
	wri:=1//(从当前这位开始写)
	i,j:=0,1
	for j<len(chars){
		if j==len(chars)-1&&chars[j]==chars[i] {
			m:=strconv.Itoa(j-i+1)
			for t:=0;t<len(m);t++{
				chars[wri]=m[t]
				wri++
			}
			break
		}
		if chars[j]!=chars[i] {
			if j-i>1 {
				s:=strconv.Itoa(j-i)
				for t:=0;t<len(s);t++{
					chars[wri]=s[t]
					wri++
				}
			}
			chars[wri]=chars[j]
			wri++
			i=j
		}
		j++
	}
	return wri
}
