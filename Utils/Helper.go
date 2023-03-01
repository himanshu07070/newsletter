package utils

import (
	"net/mail"
	"regexp"
)

func ValidateEmailAddress(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9\-]+\.[a-zA-Z]{2,}(\.[a-zA-Z]{2,})?$`)
	return emailRegex.MatchString(e)
}
