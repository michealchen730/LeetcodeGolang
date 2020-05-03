package main

import "sort"

type Ints1386 [][]int

func (s Ints1386) Len() int           { return len(s) }
func (s Ints1386) Less(i, j int) bool { return s[i][0] < s[j][0] }
func (s Ints1386) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	res := 0
	sort.Sort(Ints1386(reservedSeats))
	mp := make(map[int][]int)
	for i := 0; i < len(reservedSeats); i++ {
		mp[reservedSeats[i][0]] = append(mp[reservedSeats[i][0]], reservedSeats[i][1])
	}
	min := reservedSeats[0][0]
	max := reservedSeats[len(reservedSeats)-1][0]
	res += (min-1)*2 + (n-max)*2

	for i := min; i <= max; i++ {
		if val, ok := mp[i]; !ok {
			res += 2
			continue
		} else {
			used := make([]bool, 10)
			for _, v := range val {
				used[v-1] = true
			}
			flag0, flag1 := 0, 0
			if !(used[2] || used[3] || used[4] || used[1]) {
				res++
			} else {
				flag0 = 1
			}
			if !(used[5] || used[6] || used[7] || used[8]) {
				res++
			} else {
				flag1 = 1
			}
			if flag0 == 1 && flag1 == 1 {
				if !(used[3] || used[4] || used[5] || used[6]) {
					res++
				}
			}
		}
	}
	return res
}
