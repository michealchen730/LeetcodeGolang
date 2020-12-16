package main

import "strings"

func wordPattern(pattern string, s string) bool {
	str := strings.Fields(s)
	if len(pattern) != len(str) {
		return false
	}

	Mapping := make([]string, 26)
	mp := make(map[string]bool)

	for k, _ := range pattern {
		//此时字符串b和字符a之间没有映射关系
		if Mapping[pattern[k]-'a'] == "" {
			//假设字符串b出现过，那么说明字符串b和某个字符c之间存在映射关系
			if mp[str[k]] {
				return false
			}
			//建立映射关系
			Mapping[pattern[k]-'a'] = str[k]
			mp[str[k]] = true
		} else {
			//此时字符a和某个字符串存在映射关系
			if Mapping[pattern[k]-'a'] != str[k] {
				return false
			}
		}
	}

	return true
}
