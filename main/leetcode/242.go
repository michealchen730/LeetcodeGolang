package main

func isAnagram(s string, t string) bool {
	arr := make([]int, 26)

	for _, v := range s {
		arr[v-'a']++
	}

	for _, v := range t {
		arr[v-'a']--
		if arr[v-'a'] < 0 {
			return false
		}
	}

	for _, v := range arr {
		if v != 0 {
			return false
		}
	}

	return true

}
