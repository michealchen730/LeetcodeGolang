package main

import (
	"strings"
)

func numberToWords(num int) string {
	var (
		singles   = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
		teens     = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
		tens      = []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
		thousands = []string{"", "Thousand", "Million", "Billion"}
	)
	if num == 0 {
		return "Zero"
	}
	var sb []string
	var recursion func(n int)
	recursion = func(n int) {
		if n == 0 {
			return
		}
		if n < 10 {
			sb = append(sb, singles[n])
			return
		}
		if n < 20 {
			sb = append(sb, teens[n-10])
			return
		}
		if n < 100 {
			t := n / 10
			sb = append(sb, tens[t])
			recursion(n - t*10)
		}
		if n >= 100 {
			t := n / 100
			sb = append(sb, singles[t])
			sb = append(sb, "Hundred")
			recursion(n - t*100)
		}
	}
	for i, j := 3, int(1e9); num > 0; i-- {
		t := num / j
		recursion(t)
		if t > 0 && i > 0 {
			sb = append(sb, thousands[i])
		}
		num -= t * j
		j /= 1000
	}
	var res strings.Builder
	for k, v := range sb {
		res.WriteString(v)
		if k != len(sb)-1 {
			res.WriteString(" ")
		}
	}
	return res.String()
}
