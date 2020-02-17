package main

import (
	"strings"
)
//这题感觉没有太好的思路(后面证明了是我太垃圾了。。。)
func repeatedSubstringPattern(s string) bool {
	i:=0
	for i<len(s)/2{
		if s[i]==s[len(s)-1] {
			t1:=s[0:i+1]
			t2:=s[len(s)-i-1:len(s)]
			if strings.Compare(t1,t2)==0&&checkStrings(s,t1) {
				return true
			}
		}
		i++
	}
	return false
}

func checkStrings(s string,t string) bool{
	if len(s)%len(t)!=0{
		return false
	}else{
		nums:=len(s)/len(t)
		temp:=""
		for i:=0;i<nums;i++{
			temp=temp+t
		}
		if strings.Compare(s,temp)==0{
			return true
		}
	}
	return false
}
