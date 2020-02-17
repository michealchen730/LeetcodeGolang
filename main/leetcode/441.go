package main

func arrangeCoins(n int) int {
	//如果考虑溢出问题，减法胜过加法
	i:=1
	res:=0
	for n>=i {
		res++
		n=n-i
		i++
	}
	return res
}
