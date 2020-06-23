package main

import (
	"strconv"
	"strings"
)

type Codec struct {
	mp   map[string]string
	temp int
}

func Constructor535() Codec {
	return Codec{temp: 0, mp: make(map[string]string)}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	s := "http://tinyurl.com/"
	var res strings.Builder
	res.WriteString(s)
	res.WriteString(strconv.Itoa(this.temp))
	this.mp[res.String()] = longUrl
	this.temp++
	return res.String()
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.mp[shortUrl]
}
