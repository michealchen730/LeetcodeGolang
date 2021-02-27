package main

func mirrorReflection(p int, q int) int {
	if q == 0 {
		return 0
	}
	g := gcd(p, q)
	if p/g%2 == 0 {
		return 2
	}
	lcm := p * q / g
	//如果最小公倍数是p，那么返回1
	//如果最小公倍数不是p，那么如果最小公倍数/p为偶数，返回1
	if lcm == p || q/g%2 == 1 {
		return 1
	}
	if q/g%2 == 0 {
		return 0
	}
	return -1
}
