package main

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	mp := make(map[string]int)

	res := 0

	mp2 := make(map[byte]int)

	for i := 0; i < minSize; i++ {
		mp2[s[i]]++
	}
	if len(mp2) <= maxLetters {
		mp[s[0:minSize]]++
		res++
	}
	for j := minSize; j < len(s); j++ {
		mp2[s[j]]++
		mp2[s[j-minSize]]--
		if mp2[s[j-minSize]] == 0 {
			delete(mp2, s[j-minSize])
		}
		if len(mp2) <= maxLetters {
			mp[s[j-minSize+1:j+1]]++
			if mp[s[j-minSize+1:j+1]] > res {
				res = mp[s[j-minSize+1:j+1]]
			}
		}
	}

	return res
}
