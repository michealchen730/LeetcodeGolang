package main

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	if n1 == 0 {
		return 0
	}
	s1cnt, index, s2cnt := 0, 0, 0
	var pre_loop, in_loop []int
	mp := make(map[int][]int)
	for true {
		s1cnt++
		for k, _ := range s1 {
			if s1[k] == s2[index] {
				index++
				//可能会出现s1中包含s2的情况
				if index == len(s2) {
					s2cnt++
					index = 0
				}
			}
		}
		//没来得及出现循环节
		if s1cnt == n1 {
			return s2cnt / n2
		}
		if _, ok := mp[index]; ok {
			//出现循环节了
			s1cnt_prime, s2cnt_prime := mp[index][0], mp[index][1]    //以官方题解为例，这里应该是1和0
			pre_loop = []int{s1cnt_prime, s2cnt_prime}                //进入循环节之前的数据
			in_loop = []int{s1cnt - s1cnt_prime, s2cnt - s2cnt_prime} //循环节内的数据，这里应该是2，1
			break
		} else {
			//代表每次s1走到结尾的时候，处在s2的第几个位置上
			mp[index] = []int{s1cnt, s2cnt}
		}
	}
	ans := pre_loop[1] + (n1-pre_loop[0])/in_loop[0]*in_loop[1]
	rest := (n1 - pre_loop[0]) % in_loop[0]
	for i := 0; i < rest; i++ {
		for k, _ := range s1 {
			if s1[k] == s2[index] {
				index++
				if index == len(s2) {
					ans++
					index = 0
				}
			}
		}
	}
	return ans / n2
}
