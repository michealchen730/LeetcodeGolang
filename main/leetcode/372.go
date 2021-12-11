package main

const base = 1337

func superPow(a int, b []int) int {
	if len(b) == 0 {
		return 1
	}
	last := b[len(b)-1]
	b = b[:len(b)-1]
	p1 := myPow2(a, last)
	p2 := myPow2(superPow(a, b), 10)
	return (p1 * p2) % base
}

//func myPow2(a, b int) int{
//	a %= base
//	res:=1
//	for t:=0;t<b;t++{
//		res *= a
//		res %= base
//	}
//	return res
//}

func myPow2(a, b int) int {
	if b == 0 {
		return 1
	}
	a %= base
	if b%2 == 1 {
		return a * myPow2(a, b-1) % base
	} else {
		sub := myPow2(a, b/2)
		return sub * sub % base
	}
}
