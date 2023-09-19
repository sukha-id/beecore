package handler

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func IsValidPhoneNumber(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()
	pattern := `^\+1\d{10}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(phoneNumber)
}

func IsValidString(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()
	pattern := `^[A-Za-z0-9]+$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(phoneNumber)
}
