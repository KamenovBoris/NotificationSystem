package notification

import (
	"regexp"
	"strings"
)

func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(email)
}

func isValidBulgarianPhoneNumber(phoneNumber string) bool {
	pattern := `^\+3598[0-9]{8}$`
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(phoneNumber)
}

func isValidMessage(message string) bool {
	trimmedMessage := strings.TrimSpace(message)

	return len(trimmedMessage) > 0
}

func isValidSlackUsername(username string) bool {
	pattern := `^[a-z0-9._-]{1,21}$`
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(username)
}
