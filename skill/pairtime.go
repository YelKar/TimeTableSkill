package main

import (
	"log"
	"strings"
	"time"
)

type PairTime time.Time

func (t *PairTime) UnmarshalJSON(b []byte) error {
	timeString := strings.Trim(string(b), "\"")

	gotTime, err := time.Parse("15:04", timeString)
	if err != nil {
		log.Fatal(err)
	}
	*t = PairTime(gotTime)
	return nil
}

func (t PairTime) String() string {
	res := time.Time(t)
	return res.Format("15:04")
}
