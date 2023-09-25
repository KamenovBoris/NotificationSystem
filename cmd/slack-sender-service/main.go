package main

import (
	"fmt"
	consumer "notsys/pkg/common"
	"notsys/pkg/slack"
)

func main() {
	fmt.Printf("Slack sender service started\n")
	consumer.RunConsumerGroup([]string{"localhost:9092"}, "slack-consumer-group", []string{"slack"}, slack.SlackMessageSender{})
}
