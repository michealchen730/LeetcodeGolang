package main

import (
	"strconv"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	var bytes []byte
	//string表示被除数和除数，int表示该次除法结果的下表
	mp := make(map[string]int)

	var res strings.Builder
	if (numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0) {
		res.WriteByte('-')
	}
	numerator = abs(numerator)
	denominator = abs(denominator)
	t1 := numerator / denominator
	res.WriteString(strconv.Itoa(t1))

	numerator = numerator % denominator

	if numerator == 0 {
		return res.String()
	}
	res.WriteByte('.')
	k := 0

	for numerator != 0 {
		s := strconv.Itoa(numerator) + " " + strconv.Itoa(denominator)
		if _, ok := mp[s]; ok {
			rear := append([]byte{}, bytes[mp[s]:]...)
			bytes = append(bytes[:mp[s]], '(')
			bytes = append(bytes, rear...)
			bytes = append(bytes, ')')
			break
		} else {
			mp[s] = k
			t := numerator * 10 / denominator
			numerator = numerator * 10 % denominator
			bytes = append(bytes, byte(t+'0'))
		}
		k++
	}

	res.WriteString(string(bytes))
	return res.String()
}
