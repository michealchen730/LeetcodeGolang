package main

import "sort"

type Str0517 struct {
	strs []string
	i    int
}

//自写Sort
type S0517 []Str0517

func (s S0517) Len() int           { return len(s) }
func (s S0517) Less(i, j int) bool { return len(s[i].strs) < len(s[j].strs) }
func (s S0517) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func peopleIndexes(favoriteCompanies [][]string) []int {
	s := make([]Str0517, len(favoriteCompanies))
	for i := 0; i < len(s); i++ {
		s[i] = Str0517{strs: favoriteCompanies[i], i: i}
	}
	var res []int
	sort.Sort(S0517(s))
	res = append(res, s[len(s)-1].i)
	for i := len(s) - 2; i >= 0; i-- {
		flag1 := 0
		for j := len(s) - 1; len(s[j].strs) > len(s[i].strs); j-- {
			flag := 0
			mp := make(map[string]int)
			for _, v := range s[j].strs {
				mp[v]++
			}
			for _, v := range s[i].strs {
				if mp[v] == 0 {
					flag++
					break
				}
			}
			if flag == 0 {
				flag1 = 1
				break
			}
		}
		if flag1 == 0 {
			res = append(res, s[i].i)
		}
	}
	sort.Ints(res)
	return res
}
