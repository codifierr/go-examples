package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().AddDate(0, 0, -7*5)
	weekday := time.Duration(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	year, month, day := t.Date()
	fmt.Printf("year : %d , month : %d , day : %d", year, month, day)
	fmt.Println("")
	currentZeroDay := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	currDay := currentZeroDay.Add(-1 * (weekday - 1) * 24 * time.Hour)
	fmt.Println(currDay.UnixNano() / 1000000)
}
