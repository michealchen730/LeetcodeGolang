package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	pt1 := 0
	for pt1 < len(flowerbed) {
		if flowerbed[pt1] == 0 && (pt1 == 0 || flowerbed[pt1-1] == 0) && (pt1 == len(flowerbed)-1 || flowerbed[pt1+1] == 0) {
			flowerbed[pt1] = 1
			n--
			if n == 0 {
				return true
			}
		}
		if flowerbed[pt1] == 1 {
			pt1++
		}
		pt1++
	}
	return false
}

//太麻烦了，近似于暴力了
func canPlaceFlowers2(flowerbed []int, n int) bool {
	pt1, pt2 := -1, 0
	for pt2 != len(flowerbed) {
		if flowerbed[pt2] == 1 && pt2 != pt1 {
			if pt1 == -1 {
				n -= pt2 / 2
			} else if pt2-pt1 > 3 {
				n = n - (pt2-pt1-2)/2
			}
			if n <= 0 {
				return true
			}
			pt1 = pt2
		}
		pt2++
	}
	if pt1 == -1 {
		if (len(flowerbed)+1)/2 >= n {
			return true
		} else {
			return false
		}
	} else {
		n -= (pt2 - pt1 - 1) / 2
	}
	if n <= 0 {
		return true
	} else {
		return false
	}
}
