package email

import (
	"fmt"
)

type EmailNotification struct {
	Email   string
	Title   string
	Message string
}

func NewEmailNotification(email string, title string, message string) *EmailNotification {
	e := EmailNotification{Email: email, Title: title, Message: message}
	return &e
}

func (e *EmailNotification) Send() {
	fmt.Println("Sending Email with Title: \"" + e.Title + "\" and Body: \"" + e.Message + "\" To: " + e.Email)
}
