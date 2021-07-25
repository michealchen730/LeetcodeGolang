package main

func maximumTime(time string) string {
	//xx:xx
	res := make([]rune, 5)
	for k, v := range time {
		if v != '?' {
			res[k] = v
		} else {
			switch k {
			case 0:
				if time[1] <= '9' && time[1] >= '4' {
					res[k] = '1'
				} else {
					res[k] = '2'
				}
			case 1:
				if res[0] == '2' {
					res[k] = '3'
				} else {
					res[k] = '9'
				}
			case 3:
				res[k] = '5'
			case 4:
				res[k] = '9'
			}
		}
	}
	return string(res)
}
