package main

//感觉上像是贪心
func lemonadeChange(bills []int) bool {
	if len(bills)==0{
		return true
	}
	m5,m10:=0,0
	for i:=0;i<len(bills);i++ {
		if bills[i]==5 {
			m5++
		}else if bills[i]==10 {
			if m5>0 {
				m5--
				m10++
			}else{
				return false
			}
		}else {
			if m10>0&&m5>0 {
				m10--
				m5--
			}else if m5>=3 {
				m5=m5-3
			}else{
				return false
			}
		}
	}
	return true
}