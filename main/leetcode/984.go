package main

func strWithout3a3b(A int, B int) string {
	var res []byte
	length := A + B
	for len(res) < length {
		if A > B {
			//可以添加'a'的情况
			if !(len(res) >= 2 && res[len(res)-1] == res[len(res)-2] && res[len(res)-1] == 'a') {
				res = append(res, 'a')
				A--
			} else {
				res = append(res, 'b')
				B--
			}
		} else {
			//可以添加'a'的情况
			if !(len(res) >= 2 && res[len(res)-1] == res[len(res)-2] && res[len(res)-1] == 'b') {
				res = append(res, 'b')
				B--
			} else {
				res = append(res, 'a')
				A--
			}
		}
	}
	return string(res)

}
