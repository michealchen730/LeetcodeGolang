package main

func isLongPressedName(name string, typed string) bool {
	if len(typed) < len(name) {
		return false
	}
	if typed == name {
		return true
	}

	pt1, pt2 := 0, 0
	for pt1 < len(name) && pt2 < len(typed) {
		if name[pt1] == typed[pt2] {
			cnt1 := 1
			for pt1+1 < len(name) && name[pt1] == name[pt1+1] {
				cnt1++
				pt1++
			}
			cnt2 := 1
			for pt2+1 < len(typed) && typed[pt2] == typed[pt2+1] {
				cnt2++
				pt2++
			}
			if cnt2 < cnt1 {
				return false
			}
			pt1++
			pt2++
		} else {
			return false
		}
	}

	if pt1 == len(name) && pt2 == len(typed) {
		return true
	} else {
		return false
	}
}
