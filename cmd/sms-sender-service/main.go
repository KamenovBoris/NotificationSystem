package main

import (
	"fmt"
	consumer "notsys/pkg/common"
	"notsys/pkg/sms"
)

func main() {
	fmt.Printf("Sms sender service started\n")
	consumer.RunConsumerGroup([]string{"localhost:9092"}, "sms-consumer-group", []string{"sms"}, sms.SmsMessageSender{})
}
