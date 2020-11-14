package main

//看了题解
//这题不会
func beautifulArray(N int) []int {
	mp := make(map[int][]int)
	return split(N, mp)
}

func split(N int, mp map[int][]int) []int {
	if v, ok := mp[N]; ok {
		return v
	}

	ans := make([]int, N)
	if N == 1 {
		ans[0] = 1
	} else {
		t := 0
		for _, v := range split((N+1)/2, mp) {
			ans[t] = 2*v - 1
			t++
		}
		for _, v := range split(N/2, mp) {
			ans[t] = 2 * v
			t++
		}
	}
	mp[N] = ans
	return ans
}
