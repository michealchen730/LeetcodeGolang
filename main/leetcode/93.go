package main

//始终感觉写的太过于复杂了
func restoreIpAddresses(s string) []string {
	if len(s) < 4 {
		return nil
	}

	var res []string
	var bytes []byte

	//var dfs func(i,dot int,bytes *[]byte)

	dfs93(0, 0, s, &bytes, &res)

	return res

}

func dfs93(k, dot int, s string, bytes *[]byte, res *[]string) {
	*bytes = append(*bytes, s[k])
	//出现的点数大于3，直接返回
	if dot > 3 {
		*bytes = (*bytes)[:len(*bytes)-1]
		return
	}

	t := 1
	sum := 0
	//往前检查,当出现大于255，直接return
	for j := len(*bytes) - 1; j >= 0; j-- {
		if (*bytes)[j] != '.' {
			sum += t * int((*bytes)[j]-'0')
			if sum > 255 {
				*bytes = (*bytes)[:len(*bytes)-1]
				return
			}
		} else {
			break
		}
		t *= 10
	}

	if k == len(s)-1 {
		if dot == 3 {
			*res = append(*res, string(*bytes))
		}
		*bytes = (*bytes)[:len(*bytes)-1]
		return
	}

	*bytes = append(*bytes, '.')
	dfs93(k+1, dot+1, s, bytes, res)
	*bytes = (*bytes)[:len(*bytes)-1]

	//避免00.11.0.11以及1.01.10.11这种情况出现
	if !((s[k] == '0') && (k == 0 || (*bytes)[len(*bytes)-2] == '.')) {
		dfs93(k+1, dot, s, bytes, res)
	}

	*bytes = (*bytes)[:len(*bytes)-1]
}
