package models

type Event struct {
	ID       int    `gorm:"primaryKey"`
	DateTime string `json:"datetime"`
	Duration int    `json:"duration"`
	TimeZone string `json:"timezone"`
	FileUrl  string `json:"fileurl"`
}
