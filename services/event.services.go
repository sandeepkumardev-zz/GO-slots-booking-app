package services

import (
	"fmt"
	"slot/config"
	"slot/models"
	"slot/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Data struct {
	Date       string
	Start_Time string
	End_Time   string
	TimeZone   string
}

type AvlSlots struct {
	Time     string `json:"time"`
	IsBooked bool   `json:"is_booked"`
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func AvailableSlots(date string, timezone string) *Response {
	var events []*models.Event

	result := config.SlotDB.Order("id desc").Find(&events)

	if result.Error != nil {
		return &Response{Message: "Something went wrong!", Data: nil, Success: false}
	}

	var NewList []AvlSlots
	var BookedSlot []string

	for _, event := range events {
		// convert time string with user timezone
		cnvtTimeZone, _ := utils.ConvertTimeString(event, timezone)
		newDate, newTime := utils.SplitTime(cnvtTimeZone)

		if newDate == date {
			for _, slot := range config.TimeSlots {
				if slot == newTime {
					BookedSlot = append(BookedSlot, slot)
				}
			}
		}
	}

	for _, slot := range config.TimeSlots {
		if contains(BookedSlot, slot) {
			NewList = append(NewList, AvlSlots{slot, true})
		} else {
			NewList = append(NewList, AvlSlots{slot, false})
		}
	}

	return &Response{Message: "Successfully fetched list of events.", Data: NewList, Success: true}
}

func BookedSlots(timezone string, pageNo int) *Response {
	var events []*models.Event

	var limit = 5
	var offset = (pageNo - 1) * limit
	result := config.SlotDB.Order("date_time desc").Limit(limit).Offset(offset).Find(&events)

	if result.Error != nil {
		return &Response{Message: "Something went wrong!", Data: nil, Success: false}
	}

	if result.RowsAffected == 0 {
		return &Response{Message: "No events found with this id!", Data: nil, Success: false}
	}

	// new slice for each event
	var NewList []Data

	for _, event := range events {
		// convert time string with user timezone
		cnvtTimeZone, timeFormat := utils.ConvertTimeString(event, timezone)
		addDurationTimeZone := timeFormat.Add(time.Minute * 30).Format(time.RFC3339)

		date, start_time := utils.SplitTime(cnvtTimeZone)
		_, end_time := utils.SplitTime(addDurationTimeZone)

		NewList = append(NewList, Data{date, start_time, end_time, timezone})
	}

	return &Response{Message: "Successfully fetched list of events.", Data: NewList, Success: true}
}

type CTdata struct {
	EventId int `json:"eventId"`
}

func CreateEvent(event *models.Event) *Response {
	cnvtTimeZone, _ := utils.ConvertTimeString(event, config.TimeZone)
	_, start_time := utils.SplitTime(cnvtTimeZone)

	if !contains(config.TimeSlots, start_time) {
		return &Response{Message: "Time is not suitable for Event Host.", Data: nil, Success: false}
	}

	//check event exists or not!
	result := config.SlotDB.Where("date_time = ?", event.DateTime).First(&event)
	if result.RowsAffected != 0 {
		return &Response{Message: "Event already exists!", Data: nil, Success: false}
	}

	if err := config.SlotDB.Create(event).Error; err != nil {
		return &Response{Message: "Event creation failed!", Data: nil, Success: false}
	}

	return &Response{Message: "Successfully created event.", Data: &CTdata{EventId: event.ID}, Success: true}
}

type ResultWithFile struct {
	Event interface{} `json:"event"`
	Files interface{} `json:"files"`
}

func GetOneEvent(id string) *Response {
	var events []*models.Event

	result := config.SlotDB.Preload("Files").Where("id = ?", id).First(&events)
	if result.RowsAffected == 0 {
		return &Response{Message: "No events found with this id!", Data: nil, Success: false}
	}

	return &Response{Message: "Successfully fetched event.", Data: events[0], Success: true}
}

func UpdateEvent(id int, event *models.Event) *Response {
	event.ID = id
	result := config.SlotDB.Model(&event).Where("id = ?", id).Update(&event)

	if result.RowsAffected == 0 {
		return &Response{Message: "No events found with this id!", Data: nil, Success: false}
	}

	return &Response{Message: "Update successfully.", Data: nil, Success: true}
}

func DeleteEvent(id int, event *models.Event) *Response {
	result := config.SlotDB.Where("id = ?", id).Delete(&event)

	if result.RowsAffected == 0 {
		return &Response{Message: "No events found with this id!", Data: nil, Success: false}
	}

	return &Response{Message: "Deleted successfully.", Data: nil, Success: true}
}

func UpdateEventUrl(id int, url string) error {
	var file models.File

	file.EventId = id
	file.FileUrl = url

	if err := config.SlotDB.Create(&file).Error; err != nil {
		return err
	}

	fmt.Println(file)
	return nil
}

type FileResult struct {
	Url string `json:"url"`
}

func UploadFile(id int, ctx *gin.Context) *Response {
	var event models.Event

	result := config.SlotDB.Where("id = ?", id).First(&event)

	if result.RowsAffected == 0 {
		return &Response{Message: "No events found with this id!", Data: nil, Success: false}
	}

	file, handler, fileErr := ctx.Request.FormFile("myFile")

	if fileErr != nil {
		return &Response{Message: fileErr.Error(), Data: nil, Success: false}
	}
	defer file.Close()

	var UrlChan = make(chan string)
	var ErrChan = make(chan string)

	// upload the file, handler
	go utils.UploadToCloud(file, handler.Filename, UrlChan, ErrChan)

	select {
	case err := <-ErrChan:
		return &Response{Message: err, Data: nil, Success: false}
	case url := <-UrlChan:
		// update the File Url
		err := UpdateEventUrl(id, url)
		if err != nil {
			return &Response{Message: err.Error(), Data: nil, Success: false}
		}
		return &Response{Message: "Successfully Uploaded File.", Data: &FileResult{Url: url}, Success: true}
	}
}
