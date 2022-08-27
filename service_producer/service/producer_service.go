package service

import (
	"fmt"
	"log"
	"time"

	eventProducers "github.com/marcelobiao/poc-kafka/events/producers"
	"github.com/marcelobiao/poc-kafka/events/topics"
)

type ExampleProducerService struct {
	Producer eventProducers.IEventProducer
}

func GetExampleProducer(producer eventProducers.IEventProducer) (exampleProducerService ExampleProducerService) {
	exampleProducerService.Producer = producer
	return
}

func (e *ExampleProducerService) Run() (err error) {
	for {
		err = e.Producer.SendEvent(&topics.Example{Value: "teste"}, topics.TopicHeader{DbName: "1"})
		if err != nil {
			log.Fatalln(fmt.Errorf("send kafka message error: %s", err.Error()))
		}
		time.Sleep(time.Second * 5)
	}
}
