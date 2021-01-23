package main

func findTheDifference(s string, t string) byte {
	mp := make(map[byte]int)
	for k, _ := range t {
		mp[t[k]]++
	}
	for k, _ := range s {
		mp[s[k]]--
		if mp[s[k]] == 0 {
			delete(mp, s[k])
		}
	}
	for k, _ := range mp {
		return k
	}
	return 'n'
}

func findTheDifference2(s string, t string) byte {
	arr := make([]int, 26)
	for _, v := range t {
		arr[v-'a']++
	}
	for _, v := range s {
		arr[v-'a']--
	}
	for k, v := range arr {
		if v == 1 {
			return byte('a' + k)
		}
	}
	return 'e'
}

func findTheDifference3(s string, t string) byte {
	arr := make([]int, 26)
	for _, v := range s {
		arr[v-'a']--
	}
	for _, v := range t {
		arr[v-'a']++
		if arr[v-'a'] == 1 {
			return byte(v)
		}
	}
	return 'e'
}
