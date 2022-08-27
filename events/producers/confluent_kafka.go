package eventProducers

import (
	"encoding/json"
	"time"

	"github.com/marcelobiao/poc-kafka/events/topics"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ConfluentEventConsumer struct {
	Config   *kafka.ConfigMap
	Producer *kafka.Producer
	Chan     chan kafka.Event
}

func GetConfluentEventConsumer(config *kafka.ConfigMap, ch chan kafka.Event) (confluentEventConsumer ConfluentEventConsumer, err error) {
	producer, err := kafka.NewProducer(config)
	confluentEventConsumer.Config = config
	confluentEventConsumer.Producer = producer
	confluentEventConsumer.Chan = ch

	return
}

func (c *ConfluentEventConsumer) SendEvent(topic topics.ITopic, header topics.TopicHeader) (err error) {
	topicName := topic.GetTopicName()
	message, err := json.Marshal(topic.GetPayloadMessage())
	if err != nil {
		return Err
	}
	key, err := json.Marshal(header.DbName)
	if err != nil {
		return Err
	}

	kafkaMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topicName,
			Partition: kafka.PartitionAny,
		},
		Value:     message,
		Key:       key,
		Timestamp: time.Now(),
		//TODO: Add Headers: ,
	}
	err = c.Producer.Produce(kafkaMessage, c.Chan)
	return
}
