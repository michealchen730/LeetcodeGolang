package main

import "strconv"

func numDecodings(s string) int {
	if len(s) == 0 {
		return len(s)
	}
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	arr := make([]int, len(s))
	arr[0] = 1
	if s[1] == '0' {
		if s[0] == '1' || s[0] == '2' {
			arr[1] = 1
		}
	} else {
		temp, _ := strconv.Atoi(s[0:2])
		if 0 < temp && temp < 27 {
			arr[1] = 2
		} else {
			arr[1] = 1
		}
	}

	for i := 2; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				arr[i] = arr[i-2]
				continue
			} else {
				return 0
			}
		}
		if s[i-1] == '0' {
			arr[i] = arr[i-1]
			continue
		}
		t, _ := strconv.Atoi(s[i-1 : i+1])
		if 0 < t && t < 27 {
			arr[i] = arr[i-1] + arr[i-2]
		} else {
			arr[i] = arr[i-1]
		}
	}
	return arr[len(arr)-1]
}
