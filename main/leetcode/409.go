package main

func longestPalindrome(s string) int {
	if len(s)<=1 {
		return len(s)
	}
	temp:=make([]int,52)
	for _,v:=range s{
		if v>=97 {
			temp[v-97]++
		}else{
			temp[v-39]++
		}
	}
	f,sum:=0,0
	for i:=0; i< len(temp); i++ {
		sum=sum+temp[i]
		if f==0 {
			if temp[i]%2!=0 {
				f=1
			}
		}else{
			if temp[i]%2!=0 {
				sum--
			}
		}
	}
	return sum
}