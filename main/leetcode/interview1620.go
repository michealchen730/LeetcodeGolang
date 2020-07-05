package main

func getValidT9Words(num string, words []string) []string {
	mp := make(map[byte]int)
	for i := 0; i < 18; i++ {
		mp[byte(i+'a')] = i/3 + 2
	}
	mp['s'] = 7
	mp['t'] = 8
	mp['u'] = 8
	mp['v'] = 8
	mp['w'] = 9
	mp['x'] = 9
	mp['y'] = 9
	mp['z'] = 9
	var res []string
	for _, v := range words {
		if len(v) == len(num) {
			flag := true
			for k, _ := range v {
				if byte(mp[v[k]]+'0') != num[k] {
					flag = false
					break
				}
			}
			if flag {
				res = append(res, v)
			}
		}
	}
	return res
}
