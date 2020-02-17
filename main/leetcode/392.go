package main

func isSubsequence(s string, t string) bool {
	if len(s)==0 {
		return true
	}
	if len(t)==0 {
		return false
	}
	i,j:=0,0
	for i<len(s)&&j<len(t){
		if s[i]==t[j] {
			i++
		}
		j++
	}
	if i>=len(s) {
		return true
	}else {
		return false
	}
}