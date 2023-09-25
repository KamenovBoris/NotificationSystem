package notification

import (
	"fmt"
	"log"
	"notsys/pkg/common"

	"github.com/IBM/sarama"
)

func RunProducer(topic string, rawMessage []byte) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{common.KafkaUrl}, config)
	if err != nil {
		log.Printf("Error creating Kafka producer: %v", err)
		return err
	}
	defer producer.Close()

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(rawMessage),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Printf("Error sending message: %v", err)
		return err
	}

	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	return nil
}
