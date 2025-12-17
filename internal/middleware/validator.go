package middleware

import (
	"time"

	"age-calculator/internal/models"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func init() {
	Validate.RegisterValidation("pastdate", func(fl validator.FieldLevel) bool {
		dob, ok := fl.Field().Interface().(models.DateOnly)
		if !ok {
			return false
		}
		return time.Time(dob).Before(time.Now())
	})
}
