package slack

import (
	"fmt"
)

type SlackNotification struct {
	Recipient string
	Message   string
}

func (s *SlackNotification) Send() {
	fmt.Println("Sending Slack notification \"" + s.Message + "\" To: " + s.Recipient)
}

func NewSlackNotification(recipient string, message string) *SlackNotification {
	s := SlackNotification{Recipient: recipient, Message: message}
	return &s
}
