package main

func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
	}
	if sum%3 == 0 {
		temp, k, flag1, flag2 := 0, 0, false, false
		for k < len(A) {
			temp += A[k]
			if temp == sum/3 {
				if flag1 == false {
					flag1 = true
					temp = 0
				} else {
					flag2 = true
				}
				if flag1 && flag2 {
					return true
				}
			}
			k++
		}
	}
	return false
}
