package validator

import (
	"log"
	"slot/models"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

type CreateEvent struct {
	DateTime string `json:"datetime" validate:"required,datetime"`
	Duration int    `json:"duration" validate:"required,duration"`
	TimeZone string `json:"timezone" validate:"required,timezone"`
	FileUrl  string `json:"fileurl"`
}

type Validation struct{}

func (Validation) IsValid(event *models.Event) []string {
	var errors []string

	translator := en.New()
	uni := ut.New(translator, translator)

	trans, found := uni.GetTranslator("en")
	if !found {
		log.Fatal("translator not found")
	}

	v := validator.New()

	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		log.Fatal(err)
	}

	_ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	// DateTime
	_ = v.RegisterTranslation("datetime", trans, func(ut ut.Translator) error {
		return ut.Add("datetime", "{0} is not valid!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("datetime", fe.Field())
		return t
	})
	_ = v.RegisterValidation("datetime", func(fl validator.FieldLevel) bool {
		str := "2006-01-02 15:04"
		_, err := time.Parse(str, event.DateTime)
		return err == nil
	})

	//Duration
	_ = v.RegisterTranslation("duration", trans, func(ut ut.Translator) error {
		return ut.Add("duration", "{0} must be 30 minutes", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("duration", fe.Field())
		return t
	})
	_ = v.RegisterValidation("duration", func(fl validator.FieldLevel) bool {
		return fl.Field().Int() == 30
	})

	// Timezone
	_ = v.RegisterTranslation("timezone", trans, func(ut ut.Translator) error {
		return ut.Add("timezone", "{0} is not valid!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("timezone", fe.Field())
		return t
	})
	_ = v.RegisterValidation("timezone", func(fl validator.FieldLevel) bool {
		_, err := time.LoadLocation(event.TimeZone)
		return err == nil
	})

	a := CreateEvent{
		DateTime: event.DateTime,
		Duration: event.Duration,
		TimeZone: event.TimeZone,
	}

	err := v.Struct(a)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Translate(trans))
		}
	}

	return errors
}
