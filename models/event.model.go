package models

type Event struct {
	ID       int    `gorm:"primary_key"`
	DateTime string `json:"datetime"`
	Duration int    `json:"duration"`
	TimeZone string `json:"timezone"`
	Files    []File
}

type File struct {
	ID      int    `gorm:"primary_key"`
	FileUrl string `json:"fileurl"`
	EventId int
}
