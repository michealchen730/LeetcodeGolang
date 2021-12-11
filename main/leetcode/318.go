package main

func maxProduct318(words []string) int {
	arr := make([]int, len(words))
	for k, v := range words {
		arr[k] = convertToInt(v)
	}
	res := 0
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if arr[i]&arr[j] == 0 {
				res = max(res, len(words[i])*len(words[j]))
			}
		}
	}
	return res
}

func convertToInt(word string) int {
	t := 0
	for k, _ := range word {
		t |= 1 << int(word[k]-'a')
	}
	return t
}
