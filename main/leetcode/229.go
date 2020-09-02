package main

func majorityElement229(nums []int) []int {
	cand1, cand2 := 0, 0
	cnt1, cnt2 := 0, 0
	for _, v := range nums {
		if v == cand1 {
			cnt1++
			continue
		}

		if v == cand2 {
			cnt2++
			continue
		}

		if cnt1 == 0 {
			cand1 = v
			cnt1++
			continue
		}

		if cnt2 == 0 {
			cand2 = v
			cnt2++
			continue
		}

		cnt2--
		cnt1--
	}

	cnt1, cnt2 = 0, 0

	for _, v := range nums {
		if v == cand1 {
			cnt1++
		}
		if v == cand2 {
			cnt2++
		}
	}

	var res []int

	if cnt1 > len(nums)/3 {
		res = append(res, cand1)
	}
	if cnt2 > len(nums)/3 && cand2 != cand1 {
		res = append(res, cand2)
	}

	return res
}
