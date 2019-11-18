package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

type nsqProducer struct {
	*nsq.Producer
}

func initNsqProducer(addr string) (*nsqProducer, error) {
	fmt.Printf("init nsq producer address:%s\n", addr)
	producer, err := nsq.NewProducer(addr, nsq.NewConfig())
	if err != nil {
		return nil, err
	}
	return &nsqProducer{producer}, nil
}

func (np *nsqProducer) publish(topic, message string) error {
	err := np.Publish(topic, []byte(message))
	if err != nil {
		log.Fatal("nsq publish error:", err)
		return err
	}
	return nil
}

func main() {
	producer, err := initNsqProducer("localhost:4150")
	if err != nil {
		log.Fatal(err)
	}
	for {
		err = producer.publish("test", "testMessage")
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}

}
