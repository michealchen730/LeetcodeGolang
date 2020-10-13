package main

func smallestSubsequence(text string) string {
	var stack []byte

	s := text

	arr := make([]int, 26)

	for k, _ := range s {
		arr[s[k]-'a'] = k + 1
	}

	mp := make(map[byte]int)

	for k, _ := range s {
		//栈中有重复元素，那么直接跳过后续步骤
		if mp[s[k]] != 0 {
			continue
		}

		//判断栈顶元素出栈
		for len(stack) > 0 {
			if stack[len(stack)-1] > s[k] && arr[stack[len(stack)-1]-'a'] > k {
				mp[stack[len(stack)-1]]--
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}

		stack = append(stack, s[k])
		mp[s[k]]++
	}

	return string(stack)

}
