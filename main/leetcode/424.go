package main

func characterReplacement(s string, k int) int {
	if k >= len(s) {
		return len(s)
	}

	left, right := 0, k-1

	res := k

	cnt := make([]int, 26)
	for i := 0; i < k; i++ {
		cnt[s[i]-'A']++
	}
	mx := 0
	for _, v := range cnt {
		mx = max(mx, v)
	}
	rmn := k - (right - left + 1 - mx)

	for right < len(s) {
		if right == len(s)-1 {
			res = max(res, right-left+1)
			break
		}

		if rmn > 0 {
			right++
			cnt[s[right]-'A']++
			rmn--

			if rmn == 0 {
				for _, v := range cnt {
					mx = max(v, mx)
				}
				rmn = k - (right - left + 1 - mx)
			}

			continue
		}

		if rmn == 0 {
			res = max(res, right-left+1)
			right++
			cnt[s[right]-'A']++
			//如果该元素是当前窗口中最多的元素
			if cnt[s[right]-'A'] > mx {
				mx = cnt[s[right]-'A']
			} else {
				rmn--

				for rmn < 0 {
					cnt[s[left]-'A']--
					left++
					for _, v := range cnt {
						mx = max(mx, v)
					}
					rmn = k - (right - left + 1 - mx)
				}

			}
			continue
		}

	}

	return res
}
