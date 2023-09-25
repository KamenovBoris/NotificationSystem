package notification

import (
	"encoding/json"
	"log"
	"net/http"
	"notsys/pkg/email"
	"notsys/pkg/slack"
	"notsys/pkg/sms"
)

func EmailNotificationHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	message := r.FormValue("Message")
	receiverEmail := r.FormValue("Email")
	title := r.FormValue("Title")
	if !isValidEmail(receiverEmail) {
		http.Error(w, "Invalid Email", http.StatusBadRequest)
		return
	}
	queueNotification(email.NewEmailNotification(receiverEmail, title, message), "email", w)
}

func SmsNotificationHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	message := r.FormValue("Message")
	phoneNumber := r.FormValue("PhoneNumber")
	if !isValidBulgarianPhoneNumber(phoneNumber) {
		http.Error(w, "Invalid Bulgarian Phone Number", http.StatusBadRequest)
		return
	}
	queueNotification(sms.NewSmsNotification(phoneNumber, message), "sms", w)
}

func SlackNotificationHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	message := r.FormValue("Message")
	recipient := r.FormValue("Recipient")
	if !isValidSlackUsername(recipient) {
		http.Error(w, "Invalid Slack username", http.StatusBadRequest)
		return
	}
	if !isValidMessage(message) {
		http.Error(w, "Invalid Message! The message should contain at least one non white-space character.", http.StatusBadRequest)
		return
	}
	queueNotification(slack.NewSlackNotification(recipient, message), "slack", w)
}

func queueNotification(notification any, topic string, w http.ResponseWriter) {
	jsonData, err := json.Marshal(notification)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusBadRequest)
	} else {
		RunProducer(topic, jsonData)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Notification queued successfully!"))
	}
}
