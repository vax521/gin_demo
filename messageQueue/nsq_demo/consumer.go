package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {
	err := initConsumer("test", "pub", "127.0.0.1:4161")
	if err != nil {
		log.Fatal("init Consumer error")
	}
	select {}
}

type nsqHandler struct {
	nsqConsumer      *nsq.Consumer
	messagesReceived int
}

//处理消息
func (nh *nsqHandler) HandleMessage(msg *nsq.Message) error {
	nh.messagesReceived++
	fmt.Printf("receive ID:%s,addr:%s,message:%s", msg.ID, msg.NSQDAddress, string(msg.Body))
	fmt.Println()
	return nil
}

func initConsumer(topic, channel, addr string) error {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = 3 * time.Second
	c, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		log.Println("init Consumer NewConsumer error:", err)
		return err
	}

	handler := &nsqHandler{nsqConsumer: c}
	c.AddHandler(handler)

	err = c.ConnectToNSQLookupd(addr)
	if err != nil {
		log.Println("init Consumer ConnectToNSQLookupd error:", err)
		return err
	}
	return nil
}

/*
原文链接：https://blog.csdn.net/skh2015java/article/details/82747450
*/
