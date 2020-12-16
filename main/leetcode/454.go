package main

func fourSumCount(A []int, B []int, C []int, D []int) int {
	mp := make(map[int]int)
	mp2 := make(map[int]int)
	for k1, v1 := range A {
		for k2, v2 := range B {
			mp[v1+v2]++
			mp2[C[k1]+D[k2]]++
		}
	}

	res := 0

	for k, v := range mp {
		res += v * mp2[-k]
	}

	return res

}

//显然官方的要更加简洁
//我自己的执行时间花了官方的两倍那么多
//func fourSumCount(a, b, c, d []int) (ans int) {
//	countAB := map[int]int{}
//	for _, v := range a {
//		for _, w := range b {
//			countAB[v+w]++
//		}
//	}
//	for _, v := range c {
//		for _, w := range d {
//			ans += countAB[-v-w]
//		}
//	}
//	return
//}
