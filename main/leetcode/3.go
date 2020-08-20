package main

//很早以前的写法
func lengthOfLongestSubstring(s string) int {
	maxsize, tempsize := 0, 0
	if len(s) > 0 {
		maxsize = 1
	}
	for i := 0; i < len(s)-1; i++ {
		flag, j := 0, i+1
		for flag == 0 && j < len(s) {
			for t := i; t < j; t++ {
				if s[t] == s[j] {
					flag, tempsize = 1, j-i
					if tempsize > maxsize {
						maxsize = tempsize
					}
				}
			}
			if flag == 0 {
				j++
			}
		}
		tempsize = j - i
		if tempsize > maxsize {
			maxsize = tempsize
		}
	}
	return maxsize
}

//这题感觉滑动窗口很合适
func lengthOfLongestSubstring02(s string) int {
	mp := make(map[byte]int)
	res := 0
	lb, rb := -1, 0 //表示的是滑动窗口的左、右边界
	//这里给的边界条件是rb<len(s)，当rb到达边界，那么这个算法的结果也就出现了
	for ; rb < len(s); lb++ {
		if lb >= 0 {
			mp[s[lb]]--
		}
		for rb < len(s) && mp[s[rb]] == 0 {
			mp[s[rb]] += 1
			res = max(res, rb-lb)
			rb++
		}
	}
	return res
}
