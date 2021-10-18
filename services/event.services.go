package services

import (
	"slot/config"
	"slot/models"
	"slot/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

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

func AvailableSlots(date string, timezone string) ([]AvlSlots, string) {
	var events []*models.Event

	result := config.SlotDB.Order("event_id desc").Find(&events)

	if result.Error != nil {
		return nil, "Something went wrong!"
	}

	var NewList []AvlSlots
	var BookedSlot []string

	for _, event := range events {
		// convert time string with user timezone
		cnvtTimeZone, _ := utils.ConvertTimeString(event, timezone)
		newDate, newTime := utils.SplitTime(cnvtTimeZone)

		if newDate == date {
			for _, slot := range utils.TimeSlots {
				if slot == newTime {
					BookedSlot = append(BookedSlot, slot)
				}
			}
		}
	}

	for _, slot := range utils.TimeSlots {
		if contains(BookedSlot, slot) {
			NewList = append(NewList, AvlSlots{slot, true})
		} else {
			NewList = append(NewList, AvlSlots{slot, false})
		}
	}

	return NewList, ""
}

func BookedSlots(timezone string) ([]Data, string) {
	var events []*models.Event

	result := config.SlotDB.Order("event_id desc").Find(&events)

	if result.Error != nil {
		return nil, "Something went wrong!"
	}

	if result.RowsAffected == 0 {
		return nil, "No Booked Slots Found!"
	}

	// new slice for each event
	var NewList []Data

	for _, event := range events {
		// convert time string with user timezone
		cnvtTimeZone, timeFormat := utils.ConvertTimeString(event, timezone)
		// duration, _ := strconv.ParseInt(event.Duration, 10, 0)
		addDurationTimeZone := timeFormat.Add(time.Minute * 30).Format(time.RFC3339)

		date, start_time := utils.SplitTime(cnvtTimeZone)
		_, end_time := utils.SplitTime(addDurationTimeZone)

		NewList = append(NewList, Data{date, start_time, end_time, timezone})
	}

	return NewList, ""
}

type Result struct {
	Message  string `json:"message"`
	Response string `json:"response"`
}

func CreateEvent(event *models.Event) (*Result, string) {
	event.EventId = strconv.FormatInt(int64(time.Now().Nanosecond()), 10)

	//check event exists or not!
	result := config.SlotDB.Where("date_time = ?", event.DateTime).First(&event)
	if result.RowsAffected != 0 {
		return nil, "Event already exists!"
	}

	if err := config.SlotDB.Create(event).Error; err != nil {
		return nil, "Event creation failed!"
	}

	return &Result{Message: "Successfully created event.", Response: "Event Id: " + event.EventId}, ""
}

func GetOneEvent(id string) (*gorm.DB, string) {
	var events []*models.Event

	result := config.SlotDB.Where("event_id = ?", id).First(&events)

	if result.RowsAffected == 0 {
		return nil, "No events found with this id!"
	}

	return result, ""
}

func UpdateEvent(id string, event *models.Event) (string, string) {
	event.EventId = id
	result := config.SlotDB.Model(&event).Where("event_id = ?", id).Update(&event)

	if result.RowsAffected == 0 {
		return "", "No events found with this id!"
	}

	return "Update successfully.", ""
}

func DeleteEvent(id string, event *models.Event) (string, string) {
	result := config.SlotDB.Where("event_id = ?", id).Delete(&event)

	if result.RowsAffected == 0 {
		return "", "No events found with this id!"
	}

	return "Deleted successfully.", ""
}

func UpdateEventUrl(id string, url string) string {
	var event []*models.Event
	event[0].EventId = id
	event[0].FileUrl = url
	result := config.SlotDB.Model(&event).Where("event_id = ?", id).Update(&event)

	if result.RowsAffected == 0 {
		return "No events found with this id!"
	}

	return ""
}

func UploadFile(id string, ctx *gin.Context) (string, string) {
	var event []*models.Event

	result := config.SlotDB.Where("event_id = ?", id).First(&event)

	if result.RowsAffected == 0 {
		return "", "No events found with this id!"
	}

	file, handler, fileErr := ctx.Request.FormFile("myFile")

	if fileErr != nil {
		return "", "Error Retrieving the File"
	}
	defer file.Close()

	fileName := utils.CreateFileName(handler.Filename)

	go utils.UploadToCloud(file, fileName)

	select {
	case err := <-utils.ErrChan:
		return "", err
	case url := <-utils.UrlChan:
		err := UpdateEventUrl(id, url)
		if err != "" {
			return "", err
		}

		return "Successfully Uploaded File. " + url, ""
	}
}
