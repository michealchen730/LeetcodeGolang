package main

func canConstruct383(ransomNote string, magazine string) bool {
	arr := make([]int, 26)
	for k := range ransomNote {
		arr[int(ransomNote[k]-'a')]++
	}
	for k := range magazine {
		arr[int(magazine[k]-'a')]--
	}
	for _, v := range arr {
		if v > 0 {
			return false
		}
	}
	return true
}
