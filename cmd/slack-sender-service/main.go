package main

import (
	"fmt"
	"notsys/pkg/common"
	consumer "notsys/pkg/common-sender"
	"notsys/pkg/slack"
)

func main() {
	fmt.Printf("Slack sender service started\n")
	consumer.RunConsumerGroup([]string{common.KafkaUrl}, common.ConsumerGroupSlack, []string{common.ChannelSlack}, slack.SlackMessageSender{})
}
