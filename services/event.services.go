package services

import (
	"slot/config"
	"slot/models"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

func BookedSlots() (*gorm.DB, string) {
	var events []*models.Event
	result := config.DB.Order("date_time desc").Find(&events)

	if result.Error != nil {
		return nil, "Something went wrong!"
	}

	if result.RowsAffected == 0 {
		return nil, "No Booked Slots Found!"
	}

	return result, ""
}

type Result struct {
	message string
	eventId string
}

func CreateEvent(event *models.Event) (*Result, string) {
	event.EventId = strconv.FormatInt(int64(time.Now().Nanosecond()), 10)
	if err := config.DB.Create(event).Error; err != nil {
		return nil, "Event creation failed!"
	}

	return &Result{message: "Successfully created event.", eventId: event.EventId}, ""
}

func GetOneEvent(id string) (*gorm.DB, string) {
	var events []*models.Event

	result := config.DB.Where("event_id = ?", id).First(&events)

	if result.RowsAffected == 0 {
		return nil, "No events found with this id!"
	}

	return result, ""
}

func UpdateEvent(id string, event *models.Event) (string, string) {
	event.EventId = id
	result := config.DB.Model(&event).Where("event_id = ?", id).Update(&event)

	if result.RowsAffected == 0 {
		return "", "No events found with this id!"
	}

	return "Update successfully.", ""
}

func DeleteEvent(id string, event *models.Event) (string, string) {
	result := config.DB.Where("event_id = ?", id).Delete(&event)

	if result.RowsAffected == 0 {
		return "", "No events found with this id!"
	}

	return "Deleted successfully.", ""
}
