package main

type BrowserHistory struct {
	hist     []string
	homepage string
	temp     int
}

func Constructor1472(homepage string) BrowserHistory {
	return BrowserHistory{hist: []string{homepage}, homepage: homepage}
}

func (this *BrowserHistory) Visit(url string) {
	this.hist = this.hist[:this.temp+1]
	this.hist = append(this.hist, url)
	this.temp = len(this.hist) - 1
}

func (this *BrowserHistory) Back(steps int) string {
	this.temp = this.temp - min(steps, this.temp)
	return this.hist[this.temp]
}

func (this *BrowserHistory) Forward(steps int) string {
	this.temp = min(this.temp+steps, len(this.hist)-1)
	return this.hist[this.temp]
}
