package main

func isPalindrome09(x int) bool {
	target := x
	if x < 0 {
		return false
	}
	temp := 0
	for {
		if x != 0 {
			temp = temp*10 + x%10
			x = x / 10
		} else {
			break
		}
	}
	if temp == target {
		return true
	} else {
		return false
	}
}

//只需要反转一半的数字就可以了
//这是题解的思路
func isPalindrome10(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	target := 0
	for target < x {
		target = target*10 + x%10
		x /= 10
	}
	return x == target || x == target/10
}
