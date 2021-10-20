package utils

import (
	"slot/models"
	"strings"
	"time"
)

func SplitTime(cnvtTimeZone string) (string, string) {
	date := strings.Split(cnvtTimeZone, "T")

	var Time []string
	if strings.Contains(date[1], "+") {
		Time = strings.Split(date[1], "+")
	} else if strings.Contains(date[1], "-") {
		Time = strings.Split(date[1], "-")
	} else {
		Time[0] = "00:00"
	}

	return date[0], Time[0]
}

func SplitDate(dateTime string) []string {
	date := strings.Split(dateTime, " ")
	return date
}

//  00:30 - 07:30

func ConvertTimeString(event *models.Event, timezone string) (string, time.Time) {
	str := "2006-01-02 15:04"
	dbloc, _ := time.LoadLocation(event.TimeZone)
	dbTimeZone, _ := time.ParseInLocation(str, event.DateTime, dbloc)
	userloc, _ := time.LoadLocation(timezone)
	timeFormat := dbTimeZone.In(userloc)
	cnvtTimeZone := timeFormat.Format(time.RFC3339)

	return cnvtTimeZone, timeFormat
}
