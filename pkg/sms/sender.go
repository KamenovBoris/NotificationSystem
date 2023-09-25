package sms

import (
	commonSender "notsys/pkg/common-sender"

	"github.com/IBM/sarama"
)

type SmsMessageSender struct{}

func (h SmsMessageSender) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h SmsMessageSender) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h SmsMessageSender) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var smsNotifier SmsNotification
		commonSender.SendMessage(*msg, session, &smsNotifier)
	}
	return nil
}
