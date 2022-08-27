package eventProducers

import (
	"errors"

	"github.com/marcelobiao/poc-kafka/events/topics"
)

var (
	//TODO: Melhorar tratamento de erros
	Err = errors.New("error")
)

type IEventProducer interface {
	SendEvent(topic topics.ITopic, header topics.TopicHeader) error
}
