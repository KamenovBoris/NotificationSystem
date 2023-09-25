package sms

import (
	"fmt"
)

type SmsNotification struct {
	PhoneNumber string
	Message     string
}

func (s *SmsNotification) Send() {
	fmt.Println("Sending Sms \"" + s.Message + "\" To: " + s.PhoneNumber)
}

func NewSmsNotification(phoneNumber string, message string) *SmsNotification {
	s := SmsNotification{PhoneNumber: phoneNumber, Message: message}
	return &s
}
