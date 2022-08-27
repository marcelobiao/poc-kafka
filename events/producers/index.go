package eventProducers

import (
	"ellekrau/fullcycle-kafka-go/events/topics"
	"errors"
)

var (
	//TODO: Melhorar tratamento de erros
	Err = errors.New("error")
)

type IEventProducer interface {
	SendEvent(topic topics.ITopic, header topics.TopicHeader) error
}
