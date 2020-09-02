package main

func isNumber(s string) bool {
	mp := map[byte]byte{
		'0': 'N', '1': 'N', '2': 'N', '3': 'N', '4': 'N', '5': 'N', '6': 'N', '7': 'N', '8': 'N', '9': 'N', 'e': 'E', 'E': 'E', '.': 'P', '+': 'S', '-': 'S', ' ': 'C',
	}
	mp2 := map[string]map[byte]string{
		"initial": {
			'C': "initial",
			'N': "int",
			'P': "float",
			'S': "sign",
		},
		"sign": {
			'N': "int",
			'P': "float",
		},
		"int": {
			'N': "int",
			'E': "exp",
			'P': "point",
			'C': "end",
		},
		"point": {
			'N': "fraction",
			'E': "exp",
			'C': "end",
		},
		"float": {
			'N': "fraction",
		},
		"fraction": {
			'N': "fraction",
			'E': "exp",
			'C': "end",
		},
		"exp": {
			'N': "enum",
			'S': "esign",
		},
		"esign": {
			'N': "enum",
		},
		"enum": {
			'N': "enum",
			'C': "end",
		},
		"end": {
			'C': "end",
		},
	}

	p := "initial"
	for i := 0; i < len(s); i++ {
		if v, ok := mp[s[i]]; ok {
			if _, ok := mp2[p][v]; ok {
				p = mp2[p][v]
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return p == "int" || p == "fraction" || p == "enum" || p == "end" || p == "point"
}
