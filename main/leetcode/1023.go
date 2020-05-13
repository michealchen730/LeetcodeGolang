package main

import (
	"fmt"
	"regexp"
	"strings"
)

func camelMatch(queries []string, pattern string) []bool {
	var str strings.Builder
	str.WriteByte('^')
	str.WriteString("[a-z]*")
	for k, _ := range pattern {
		str.WriteByte(pattern[k])
		str.WriteString("[a-z]*")
	}
	str.WriteByte('$')
	res2 := regexp.MustCompile(str.String())
	//结算
	var arr []bool
	for _, v := range queries {
		arr = append(arr, res2.MatchString(v))
	}
	return arr
}

func main() {
	fmt.Println(camelMatch([]string{"aksvbjLiknuTzqono", "ksvjLimflkpnTzqn", "mmkasvjLiknTxzqn", "ksvjLiurknTzzqbn", "ksvsjLctikgnTzqn", "knzsvzjLiknTszqn"}, "ksvjLiknTzqn"))
}
