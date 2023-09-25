package slack

import (
	"notsys/pkg/common"

	"github.com/IBM/sarama"
)

type SlackMessageSender struct{}

func (h SlackMessageSender) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h SlackMessageSender) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h SlackMessageSender) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var slackNotifier SlackNotification
		common.SendMessage(*msg, session, &slackNotifier)
	}
	return nil
}
