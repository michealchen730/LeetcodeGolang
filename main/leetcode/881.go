package main

func numRescueBoats(people []int, limit int) int {
	p := make([]int, 30001)
	mx := 0
	for _, v := range people {
		mx = max(mx, v)
		p[v]++
	}
	ans := 0
	for mx > 0 {
		if p[mx] == 0 {
			mx--
			continue
		}
		//希望这一个循环要装满一个船
		ans += 1
		p[mx]--
		l := limit - mx
	FI:
		for l != 0 {
			t := l
			for j := l; j > 0; j-- {
				if p[j] != 0 {
					p[j] -= 1
					l = l - j
					break FI
				}
			}
			if t == l {
				break
			}
		}
	}
	return ans
}
