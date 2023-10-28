package main

import "fmt"

type TimeTable []TableDay
type TableDay []Pair

func (d *TableDay) Answer() (dayString string, err error) {
	if len(*d) == 0 {
		err = FreeDayError{
			"В этот день пар нет. " +
				"Этот день как чистый лист, " +
				"которому предстоит быть заполненным. " +
				"Ваша цель не просрать его, как это обычно происходит",
		}
		return
	}
	dayString = fmt.Sprintf("\nПары начинаются в %s", (*d)[0].Start)
	var v Pair
	var i int
	for i, v = range *d {
		dayString += fmt.Sprintf("\n%d-я пара\n%s", i+1, v.Answer())
	}
	dayString += fmt.Sprintf("\n\nОкончание в %s", v.End)
	return
}
