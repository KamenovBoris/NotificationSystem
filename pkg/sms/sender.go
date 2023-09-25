package sms

import (
	"notsys/pkg/common"

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
		common.SendMessage(*msg, session, &smsNotifier)
	}
	return nil
}
