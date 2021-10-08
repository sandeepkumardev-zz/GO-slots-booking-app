package models

import "time"

type Event struct {
	EventId  string    `json:"eventid"`
	DateTime time.Time `json:"datetime"`
	Duration string    `json:"duration"`
	TimeZone string    `json:"timezone"`
}
