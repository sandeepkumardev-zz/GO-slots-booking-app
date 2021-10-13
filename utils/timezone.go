package utils

import "strings"

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

var TimeSlots = []string{"10:00:00", "10:30:00", "11:00:00", "11:30:00", "12:00:00", "12:30:00", "13:00:00", "13:30:00", "14:00:00", "14:30:00", "15:00:00", "15:30:00", "16:00:00", "16:30:00", "17:00:00"}
