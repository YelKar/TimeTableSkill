package main

import "fmt"

type TimeTable []TableDay
type TableDay []Pair

func (d *TableDay) Answer() (dayString string) {
	dayString = fmt.Sprintf("\nПары начинаются в %s", (*d)[0].Start)
	var v Pair
	var i int
	for i, v = range *d {
		dayString += fmt.Sprintf("\n%d-я пара\n%s", i+1, v.Answer())
	}
	dayString += fmt.Sprintf("\n\nОкончание в %s", v.End)
	return
}
