package validator

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type CreateEvent struct {
	EventId  string `json:"eventid"`
	DateTime string `json:"datetime"`
	Duration string `json:"duration"`
	TimeZone string `json:"timezone"`
}

type Booking struct {
	CheckIn  string `json:"check_in" binding:"required,Bookabledate" time_format:"2006-01-02"`
	CheckOut string `json:"check_out" binding:"required,Bookabledate" time_format:"2006-01-02"`
}

var BookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}
