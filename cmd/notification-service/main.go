package main

import (
	"fmt"
	"log"
	"net/http"
	"notsys/pkg/common"
	"notsys/pkg/notification"
)

func main() {
	fmt.Printf("Notification Service started on port %s\n", common.NetworkServicePort)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/api/notify/sms", notification.SmsNotificationHandler)
	http.HandleFunc("/api/notify/email", notification.EmailNotificationHandler)
	http.HandleFunc("/api/notify/slack", notification.SlackNotificationHandler)
	if err := http.ListenAndServe(":"+common.NetworkServicePort, nil); err != nil {
		log.Fatalf(err.Error())
	}
}
