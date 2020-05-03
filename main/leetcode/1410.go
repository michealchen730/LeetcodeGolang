package main

import "strings"

func entityParser(text string) string {
	var news strings.Builder
	i := 0
	for i < len(text) {
		if text[i] != '&' {
			news.WriteByte(text[i])
		} else {
			if i+4 <= len(text) {
				t2 := text[i : i+4]
				if t2 == "&gt;" {
					news.WriteByte('>')
					i += 4
					continue
				}
				if t2 == "&lt;" {
					news.WriteByte('<')
					i += 4
					continue
				}
			}
			if i+5 <= len(text) {
				t3 := text[i : i+5]
				if t3 == "&amp;" {
					news.WriteByte('&')
					i += 5
					continue
				}
			}
			if i+6 <= len(text) {
				t1 := text[i : i+6]
				if t1 == "&quot;" {
					news.WriteString("\"")
					i += 6
					continue
				}
				if t1 == "&apos;" {
					news.WriteString("'")
					i += 6
					continue
				}
			}
			if i+7 <= len(text) {
				t4 := text[i : i+7]
				if t4 == "&frasl;" {
					news.WriteByte('/')
					i += 7
					continue
				}
			}
			news.WriteByte(text[i])
		}
		i++
	}
	return news.String()
}
