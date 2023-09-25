package email

import (
	"notsys/pkg/common"

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
		common.SendMessage(*msg, session, &emailNotifier)
	}
	return nil
}
