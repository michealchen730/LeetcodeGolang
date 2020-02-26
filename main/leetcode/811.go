package main

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	mp := make(map[string]int)
	for _, v := range cpdomains {
		flds := strings.Fields(v)
		count, _ := strconv.Atoi(flds[0])
		mp[flds[1]] += count
		temp := flds[1]
		for strings.Count(temp, ".") > 0 {
			temp = temp[strings.Index(temp, ".")+1:]
			mp[temp] += count
		}
	}
	var res []string
	for k, v := range mp {
		str := strconv.Itoa(v) + " " + k
		res = append(res, str)
	}
	return res
}
