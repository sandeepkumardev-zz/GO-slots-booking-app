package validater

type CreateEvent struct {
	EventId  string `json:"eventid"`
	DateTime string `json:"datetime"`
	Duration string `json:"duration"`
	TimeZone string `json:"timezone"`
}
