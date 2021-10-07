package main

import "fmt"

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	negative := false
	arr := make([]int, 32)
	if num < 0 {
		negative = true
		num = -num
	}
	i := len(arr) - 1
	for num > 0 {
		arr[i] = num % 2
		num /= 2
		i--
	}
	if negative {
		change := false
		for k, _ := range arr {
			if change {
				arr[31-k] ^= 1
				continue
			}
			if arr[31-k] == 1 {
				change = true
			}
		}
	}
	var res []byte
	tmp := 0
	for k, _ := range arr {
		tmp <<= 1
		tmp += arr[k]
		if (k+1)%4 == 0 {
			if tmp >= 10 {
				res = append(res, byte('a'+tmp-10))
			} else if tmp > 0 {
				res = append(res, byte('0'+tmp))
			} else {
				if len(res) != 0 {
					res = append(res, '0')
				}
			}
			tmp = 0
		}
	}
	return string(res)
}

func main() {
	fmt.Println(toHex(-1))
}
