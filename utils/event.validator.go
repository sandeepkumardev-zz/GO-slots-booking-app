package utils

import (
	"slot/models"
)

func CreateValidator(event *models.Event) string {
	if event.Duration != "30" {
		return "Duration must be 30 min."
	}

	// re := regexp.MustCompile(`^([0-9]|0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$`)
	// b := re.MatchString(event.DateTime)

	// goment.New()

	// c, d := goment.New(event.DateTime)

	// fmt.Println(b)
	return ""
}
