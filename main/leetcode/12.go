package main

func intToRoman(num int) string {
	res:=""
	for num>=1000 {
		res=res+"M"
		num=num-1000
	}
	if num>=900 {
		res=res+"CM"
		num=num-900
	}
	if num>=500 {
		res=res+"D"
		num=num-500
	}
	if num>=400 {
		res=res+"CD"
		num=num-400
	}
	for num>=100 {
		res=res+"C"
		num=num-100
	}
	if num>=90 {
		res=res+"XC"
		num=num-90
	}
	if num>=50 {
		res=res+"L"
		num=num-50
	}
	if num>=40 {
		res=res+"XL"
		num=num-40
	}
	for num>=10 {
		res=res+"X"
		num=num-10
	}
	if num>=9 {
		res=res+"IX"
		num=num-9
	}
	if num>=5 {
		res=res+"V"
		num=num-5
	}
	if num>=4 {
		res=res+"IV"
		num=num-4
	}
	for num>=1 {
		res=res+"I"
		num=num-1
	}
	return res
}
