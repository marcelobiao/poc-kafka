package main

import (
	"fmt"
	"log"

	eventProducers "github.com/marcelobiao/poc-kafka/events/producers"
	"github.com/marcelobiao/poc-kafka/service_producer/service"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	config := kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}
	ch := make(chan kafka.Event)
	go runKafkaMessageDeliveryReport(ch)
	confluentProducer, err := eventProducers.GetConfluentEventConsumer(&config, ch)
	if err != nil {
		return
	}
	producer := service.GetExampleProducer(&confluentProducer)
	err = producer.Run()
	if err != nil {
		return
	}
}

func runKafkaMessageDeliveryReport(deliveryChan chan kafka.Event) {
	for event := range deliveryChan {
		switch event.(type) {
		case *kafka.Message:
			if err := event.(*kafka.Message).TopicPartition.Error; err != nil {
				log.Fatalln(fmt.Errorf("send kafka message error: %s", err.Error()))
			}
			log.Printf("message was successfully sent to '%s'", event.(*kafka.Message).TopicPartition)
		}
	}
}
