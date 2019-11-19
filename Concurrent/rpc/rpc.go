package main

import (
	"errors"
	"fmt"
	"time"
)

func RPCClient(ch chan string, req string) (string, error) {
	//向服务器发送请求
	ch <- req
	select {
	case ack := <-ch:
		return ack, nil
	case <-time.After(time.Second): //超时
		return "", errors.New("Time out")
	}
}

func RPCServer(ch chan string) {
	for {
		//接受请求
		data := <-ch
		fmt.Println("server received:", data)
		time.Sleep(time.Second * 2)
		//向客户端反馈
		ch <- "roger"
	}
}
func main() {
	ch := make(chan string)
	go RPCServer(ch)
	recv, err := RPCClient(ch, "hi")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("client received:", recv)
	}
}
