package main

func findTheLongestSubstring(s string) int {
	pos := make([]int, 32)
	for i := 1; i < len(pos); i++ {
		pos[i] = -1
	}
	res, status := 0, 0
	for k, v := range s {
		switch v {
		case 'a':
			status ^= 1
		case 'e':
			status ^= 2
		case 'i':
			status ^= 4
		case 'o':
			status ^= 8
		case 'u':
			status ^= 16
		}
		if pos[status] >= 0 {
			res = max(res, k-pos[status]+1)
		} else {
			pos[status] = k + 1
		}
	}
	return res
}
