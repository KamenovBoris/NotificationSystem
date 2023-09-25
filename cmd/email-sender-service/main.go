package main

import (
	"fmt"
	consumer "notsys/pkg/common"
	"notsys/pkg/email"
)

func main() {
	fmt.Printf("Email sender service started\n")
	consumer.RunConsumerGroup([]string{"localhost:9092"}, "email-consumer-group", []string{"email"}, email.EmailMessageSender{})
}
