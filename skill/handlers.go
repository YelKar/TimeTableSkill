package main

import (
	"time"
)

func getDayAfterToday(after int) (string, error) {
	weekday := time.Now().Weekday()
	return getByWeekday((int(weekday) + 6 + int(after)) % 7)
}

func getByWeekday(dayNum int) (string, error) {
	table := getTable()
	return table[dayNum].Answer()
}
