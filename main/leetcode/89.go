package main

//方法使用的是镜像反射法
//简单来说就是
//n==0 a1=[0]
//n==1 将a1翻转过来，每个都^1，然后与原来的a1拼接，得到a2[0,1]
//n==2 将a2翻转过来，每个都^2，然后与原来的a2拼接，得到a3[0,1,3,2]
//n==2 将a3翻转过来，每个都^4，然后与原来的a3拼接，得到a4[0,1,3,2,6,7,5,4]

func grayCode(n int) []int {
	t := 1
	res := make([]int, 1<<n)
	for n >= 1 {
		temp := t
		for k := t; k < 2*t; k++ {
			res[k] = res[temp-1] ^ t
			temp--
		}
		t = t << 1
		n--
	}
	return res
}
