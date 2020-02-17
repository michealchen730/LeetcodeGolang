package main

func dayOfTheWeek(day int, month int, year int) string {
	res:=[]string{"Sunday","Monday","Tuesday","Wednesday","Thursday","Friday","Saturday"}
	if month==1{
		month=13
		year--
	}
	if month==2{
		month=14
		year--
	}
	return res[(day+2*month+3*(month+1)/5+year+year/4-year/100+year/400+1)%7]
}