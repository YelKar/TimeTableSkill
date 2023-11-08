package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

const (
	monday = iota
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

const (
	anyWeek = iota
	redWeek
	blueWeek
)

var weekdays = []string{
	"понедельник",
	"вторник",
	"среда",
	"четверг",
	"пятница",
	"суббота",
	"воскресенье",
}

var daysAfterToday = []string{
	"сегодня",
	"завтра",
	"послезавтра",
}

func getTable() (table TimeTable) {
	file, _ := os.Open("timetable.json")
	data, _ := io.ReadAll(file)

	err := json.Unmarshal(data, &table)
	if err != nil {
		log.Fatal(err)
	}
	return
}
