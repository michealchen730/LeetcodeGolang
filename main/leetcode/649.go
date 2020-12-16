package main

func predictPartyVictory(senate string) string {
	V := make([]bool, len(senate))
	c1, c2 := 0, 0
	for _, v := range senate {
		if v == 'R' {
			c1++
		} else {
			c2++
		}
	}

	fR, fD := 0, 0

	for c1 != 0 && c2 != 0 {
		for k, v := range senate {
			//如果该参议员丧失权力
			if V[k] {
				continue
			}
			//该参议员没有丧失权力
			if v == 'R' {
				//此时他不会丧失权力
				if fR == 0 {
					fD++
				} else {
					V[k] = true
					c1--
					fR--
				}
			} else {
				if fD == 0 {
					fR++
				} else {
					V[k] = true
					c2--
					fD--
				}
			}
		}
	}

	if c1 == 0 {
		return "Dire"
	} else {
		return "Radiant"
	}
}
