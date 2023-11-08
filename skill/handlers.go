package main

import (
	"time"
)

func getDayAfterToday(after int) (string, error) {
	weekday := time.Now().Weekday()
	day := getByWeekday((int(weekday) + 6 + after) % 7)
	return day.Answer()
}

func getByWeekdayAnswer(dayNum int) (string, error) {
	day := getByWeekday(dayNum)
	return day.Answer()
}

func getByWeekday(dayNum int) TableDay {
	table := getTable()
	return table[dayNum]
}
