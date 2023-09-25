package email

import (
	commonSender "notsys/pkg/common-sender"

	"github.com/IBM/sarama"
)

type EmailMessageSender struct{}

func (h EmailMessageSender) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h EmailMessageSender) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h EmailMessageSender) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var emailNotifier EmailNotification
		commonSender.SendMessage(*msg, session, &emailNotifier)
	}
	return nil
}
