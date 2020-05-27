package main

import (
	"sort"
	"strings"
)

func originalDigits(s string) string {
	mp := make(map[byte]int)
	for k, _ := range s {
		mp[s[k]]++
	}
	var nums []int
	for mp['z'] != 0 {
		nums = append(nums, 0)
		mp['z']--
	}
	for mp['w'] != 0 {
		nums = append(nums, 2)
		mp['w']--
	}
	for mp['u'] != 0 {
		nums = append(nums, 4)
		mp['f']--
		mp['u']--
	}
	for mp['x'] != 0 {
		nums = append(nums, 6)
		mp['s']--
		mp['i']--
		mp['x']--
	}
	for mp['g'] != 0 {
		nums = append(nums, 8)
		mp['i']--
		mp['g']--
		mp['h']--
	}
	for mp['h'] != 0 {
		nums = append(nums, 3)
		mp['h']--
	}
	for mp['f'] != 0 {
		nums = append(nums, 5)
		mp['f']--
		mp['i']--
	}
	for mp['s'] != 0 {
		nums = append(nums, 7)
		mp['s']--
		mp['n']--
	}
	for mp['i'] != 0 {
		nums = append(nums, 9)
		mp['i']--
		mp['n'] -= 2
	}
	for mp['n'] != 0 {
		nums = append(nums, 1)
		mp['n']--
	}
	sort.Ints(nums)
	var res strings.Builder
	for _, v := range nums {
		res.WriteByte(byte(48 + v))
	}
	return res.String()
}
