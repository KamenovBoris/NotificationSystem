package main

import (
	"fmt"
	"notsys/pkg/common"
	consumer "notsys/pkg/common-sender"
	"notsys/pkg/email"
)

func main() {
	fmt.Printf("Email sender service started\n")
	consumer.RunConsumerGroup([]string{common.KafkaUrl}, common.ConsumerGroupEmail, []string{common.ChannelEmail}, email.EmailMessageSender{})
}
