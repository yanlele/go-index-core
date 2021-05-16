package date2week

import (
	"time"
)

func Date2Week(params ...time.Time) int {
	date := time.Now()

	for index, param := range params {
		if index == 0 {
			date = param
		}
	}

	year := date.Year()
	month := int(date.Month()) //time.Now().Month().String()
	day := date.Day()

	var weekday = [7]int{7, 1, 2, 3, 4, 5, 6}
	var y, m, c int
	if month >= 3 {
		m = month
		y = year % 100
		c = year / 100
	} else {
		m = month + 12
		y = (year - 1) % 100
		c = (year - 1) / 100
	}
	week := y + (y / 4) + (c / 4) - 2*c + ((26 * (m + 1)) / 10) + day - 1
	if week < 0 {
		week = 7 - (-week)%7
	} else {
		week = week % 7
	}
	whichWeek := week
	return weekday[whichWeek]
}
