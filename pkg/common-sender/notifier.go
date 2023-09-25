package commonSender

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
)

type Notifier interface {
	Send()
}

func SendMessage[T Notifier](msg sarama.ConsumerMessage, session sarama.ConsumerGroupSession, notifier T) {
	if err := json.Unmarshal(msg.Value, &notifier); err != nil {
		log.Printf("Error unmarshaling notification JSON: %v", err)
	} else {
		notifier.Send()
		session.MarkMessage(&msg, "")
	}
}
