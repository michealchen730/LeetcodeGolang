package main

func romanToInt(s string) int{
	res:=0
	r2i:= map[byte]int{'I': 1, 'V': 5,'X':10,'L':50,'C':100,'D':500,'M':1000}
	for i:=0; i<len(s);i++ {
		if i!=len(s)-1{
			if r2i[s[i]]<r2i[s[i+1]]{
				res-=r2i[s[i]]
			}else{
				res+=r2i[s[i]]
			}
		}else{
			res+=r2i[s[i]]
		}
	}
	return res
}