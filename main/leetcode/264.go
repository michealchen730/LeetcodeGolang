package main

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	l2, l3, l5 := []int{2}, []int{3}, []int{5}
	mp := make(map[int]bool)
	i, j, k, t, t1, t2, t3 := 0, 0, 0, 2, 0, 0, 0
	for t != n {
		if l2[i] < l3[j] && l2[i] < l5[k] {
			t1, t2, t3 = l2[i]*2, l2[i]*3, l2[i]*5
			i++
		} else if l3[j] < l2[i] && l3[j] < l5[k] {
			t1, t2, t3 = l3[j]*2, l3[j]*3, l3[j]*5
			j++
		} else {
			t1, t2, t3 = l5[k]*2, l5[k]*3, l5[k]*5
			k++
		}
		if !mp[t1] {
			l2 = append(l2, t1)
			mp[t1] = true
		}
		if !mp[t2] {
			l3 = append(l3, t2)
			mp[t2] = true
		}
		if !mp[t3] {
			l5 = append(l5, t3)
			mp[t3] = true
		}
		t++
	}
	return min(min(l2[i], l3[j]), l5[k])
}
