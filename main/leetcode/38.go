package main

import (
	"strconv"
	"strings"
)

//很早以前的写法，从java转过来的
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	} else if n == 2 {
		return "11"
	} else {
		temp := countAndSay(n - 1)

		bytes := []byte(temp)
		var sb strings.Builder

		flag, i := 1, 1

		for i < len(bytes) {
			if bytes[i] == bytes[i-1] {
				flag++
			} else {
				sb.WriteString(strconv.Itoa(flag))
				sb.WriteByte(bytes[i-1])
				flag = 1
			}
			i++
		}
		sb.WriteString(strconv.Itoa(flag))
		sb.WriteByte(bytes[i-1])

		return sb.String()
	}

}

func countAndSay38(n int) string {
	bytes := []byte{'1'}
	for n > 1 {
		count := 0
		var temp []byte
		for i := 0; i < len(bytes); i++ {
			if i == 0 || bytes[i] == bytes[i-1] {
				count++
				continue
			}
			if bytes[i-1] != bytes[i] {
				temp = append(temp, []byte(strconv.Itoa(count))...)
				temp = append(temp, bytes[i-1])
				count = 1
			}
		}
		temp = append(temp, []byte(strconv.Itoa(count))...)
		temp = append(temp, bytes[len(bytes)-1])

		bytes = temp
		n--
	}
	return string(bytes)
}
