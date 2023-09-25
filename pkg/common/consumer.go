package common

import (
	"context"
	"log"

	"github.com/IBM/sarama"
)

type MessageSender interface {
	Setup(sarama.ConsumerGroupSession) error
	Cleanup(sarama.ConsumerGroupSession) error
	ConsumeClaim(sarama.ConsumerGroupSession, sarama.ConsumerGroupClaim) error
}

func RunConsumerGroup(brokers []string, group string, topics []string, sender MessageSender) {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumerGroup(brokers, group, config)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer group: %v", err)
	}
	defer consumer.Close()

	for {
		if err := consumer.Consume(context.Background(), topics, sender); err != nil {
			log.Fatalf("Error consuming messages: %v", err)
		}
	}
}
