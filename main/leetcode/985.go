package main

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	sum := 0
	for _, v := range A {
		if v%2 == 0 {
			sum += v
		}
	}
	var res []int
	for _, v := range queries {
		val, ind := v[0], v[1]
		flag1 := A[ind]%2 == 0
		flag2 := val%2 == 0
		A[ind] += val
		if flag1 && !flag2 { //even odd
			sum -= A[ind] - val
		} else if !flag1 && !flag2 { //odd odd
			sum += A[ind]
		} else if flag1 && flag2 { //even even
			sum += val
		}
		res = append(res, sum)
	}
	return res
}
