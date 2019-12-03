package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DataEvent struct {
	Data  interface{}
	Topic string
}

type dataChan chan DataEvent

type DataChanSlice []dataChan

//事件总线
type EventBus struct {
	subscribers map[string]DataChanSlice
	rm          sync.RWMutex
}

func (eb *EventBus) Subscribe(topic string, ch dataChan) {
	eb.rm.Lock()
	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append([]dataChan{}, ch)
	}
	eb.rm.Unlock()
}

func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.rm.RLock()
	if chans, found := eb.subscribers[topic]; found {
		newChans := append(DataChanSlice{}, chans...)
		go func(data DataEvent, dataChans DataChanSlice) {
			for _, ch := range dataChans {
				ch <- data
			}
		}(DataEvent{Topic: topic, Data: data}, newChans)
	}
	eb.rm.RUnlock()
}

var eb = &EventBus{
	subscribers: map[string]DataChanSlice{},
}

func publishTo(topic string, data string) {

	eb.Publish(topic, data)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

}

func printDataEvent(ch string, data DataEvent) {
	fmt.Printf("Channel:%s,Topic:%s,DataEvent:%v;\n", ch, data.Topic, data.Data)
}

func main() {
	chan1 := make(chan DataEvent)
	chan2 := make(chan DataEvent)
	chan3 := make(chan DataEvent)

	eb.Subscribe("topic1", chan1)
	eb.Subscribe("topic2", chan1)
	eb.Subscribe("topic2", chan2)
	eb.Subscribe("topic3", chan3)
	go publishTo("topic1", "hi topic1")
	go publishTo("topic2", "hi topic2")
	go publishTo("topic3", "hi topic3")

	for {
		select {
		case d := <-chan1:
			{
				go printDataEvent("ch1", d)
			}
		case d := <-chan2:
			{
				go printDataEvent("ch2", d)
			}
		case d := <-chan3:
			{
				go printDataEvent("ch3", d)
			}
		}
	}
	time.Sleep(time.Second)

}
