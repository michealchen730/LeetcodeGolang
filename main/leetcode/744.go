package main

func nextGreatestLetter(letters []byte, target byte) byte {
	//target大于等于letters内所有字符
	if int(target)-int(letters[len(letters)-1]) >= 0 {
		return letters[0]
	}
	low, high := 0, len(letters)-1
	for low < high {
		mid := low + (high-low)/2
		if int(letters[mid])-int(target) > 0 {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return letters[low]
}
