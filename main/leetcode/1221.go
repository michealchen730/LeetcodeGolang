package main

func balancedStringSplit(s string) int {
	i:=0
	count:=0
	sum:=0
	for i<=len(s)-1{
		if s[i]=='L'{
			sum++
		}else{
			sum--
		}
		if sum==0 {
			count++
		}
		i++
	}
	return count
}

