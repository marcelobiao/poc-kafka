package eventProducers

import (
	"encoding/json"

	"github.com/marcelobiao/poc-kafka/events/topics"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ConfluentEventProducer struct {
	Config   *kafka.ConfigMap
	Producer *kafka.Producer
	Chan     chan kafka.Event
}

func NewConfluentEventProducer(config *kafka.ConfigMap, ch chan kafka.Event) (confluentEventProducer ConfluentEventProducer, err error) {
	producer, err := kafka.NewProducer(config)
	confluentEventProducer.Config = config
	confluentEventProducer.Producer = producer
	confluentEventProducer.Chan = ch

	return
}

func (c *ConfluentEventProducer) SendEvent(topic topics.ITopic, header topics.TopicHeader) (err error) {
	topicName := topic.GetTopicName()
	message, err := json.Marshal(topic.GetPayloadMessage())
	if err != nil {
		return Err
	}

	kafkaMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topicName,
			Partition: kafka.PartitionAny,
		},
		Value: message,
		//TODO: Add Headers: ,
	}
	err = c.Producer.Produce(kafkaMessage, c.Chan)
	return
}
