package main

import (
	"fmt"
	"log"
	"net/http"
	"notsys/pkg/notification"
)

const port = "8100"

func main() {
	fmt.Printf("Notification Service started on port %s\n", port)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/api/notify/sms", notification.SmsNotificationHandler)
	http.HandleFunc("/api/notify/email", notification.EmailNotificationHandler)
	http.HandleFunc("/api/notify/slack", notification.SlackNotificationHandler)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		errorMessage := err.Error()
		log.Fatalf(errorMessage)
		fmt.Printf(errorMessage)
	}
}
