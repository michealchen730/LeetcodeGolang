package main

import "math"

//这是道我完全不会的题。。。
func divide(dividend int, divisor int) int {
	//sign用来判断最终的结果的正负
	//这里sign==0是表示最终结果为正
	sign := 1
	if (dividend >= 0 && divisor > 0) || (dividend <= 0 && divisor < 0) {
		sign = 0
	}
	//有了sign的判断，为了方便处理，就可以对除数和被除数进行abs操作
	dividend, divisor = abs(dividend), abs(divisor)
	res, p, cmp := 0, 1, divisor

	//除数小于被除数，那么自然为0
	if dividend < cmp {
		return 0
	}
	// 第一步，例子：60/8，首先移位是要得出4，4*8==32
	//这一步最终得到cmp=32,p=4
	for (cmp << 1) < dividend {
		cmp = cmp << 1
		p = p << 1
	}

	// 例子：60需要拿掉32，以后剩余28，还是大于8
	for dividend >= divisor {
		dividend -= cmp
		res += p
		//此时cmp>a(这是一定的，如果32没有大于28，那么第一步那里一定可以继续执行下去)
		for cmp > dividend {
			//此时将cmp往回挪
			cmp = cmp >> 1
			p = p >> 1
		}
	}
	//60->28->12->4
	//32->16-> 8->4
	// 4-> 2-> 1

	if sign == 1 {
		if -res < math.MinInt32 {
			return math.MaxInt32
		}
		return -res
	} else {
		if res > math.MaxInt32 {
			return math.MaxInt32
		}
		return res
	}
}
