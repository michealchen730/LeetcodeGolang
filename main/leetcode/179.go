package main

import (
	"sort"
	"strconv"
	"strings"
)

type Sort179 []int

func (s Sort179) Len() int { return len(s) }
func (s Sort179) Less(i, j int) bool {
	s1 := strconv.Itoa(s[i])
	s2 := strconv.Itoa(s[j])
	t1, t2 := s1+s2, s2+s1

	for i := 0; i < len(t1); i++ {
		if t1[i] > t2[i] {
			return true
		} else if t1[i] < t2[i] {
			return false
		}
	}
	return true
}
func (s Sort179) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func largestNumber179(nums []int) string {
	sort.Sort(Sort179(nums))
	if nums[0] == 0 && nums[len(nums)-1] == 0 {
		return "0"
	}

	var res strings.Builder
	for i := 0; i < len(nums); i++ {
		res.WriteString(strconv.Itoa(nums[i]))
	}
	return res.String()
}
