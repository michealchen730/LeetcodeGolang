package main

func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}
	if len(s) == k {
		return true
	}
	arr := make([]int, 26)
	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
	}
	nums := 0
	for _, v := range arr {
		if v%2 != 0 {
			nums++
		}
	}
	if nums > k {
		return false
	}
	if nums == k {
		return true
	}
	if nums < k {
		k = k - nums
		if len(s)-nums >= k {
			return true
		} else {
			return false
		}
	}
	return false
}
