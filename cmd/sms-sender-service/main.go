package main

import (
	"fmt"
	"notsys/pkg/common"
	consumer "notsys/pkg/common-sender"
	"notsys/pkg/sms"
)

func main() {
	fmt.Printf("Sms sender service started\n")
	consumer.RunConsumerGroup([]string{common.KafkaUrl}, common.ConsumerGroupSms, []string{common.ChannelSms}, sms.SmsMessageSender{})
}
