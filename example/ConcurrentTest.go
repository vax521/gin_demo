package main

import (
	"fmt"
	"math/rand"
	"time"
)

//数据生产者
func producer(header string, channel chan<- string) {
	//无限循环，不停的生产数据
	for {
		// 将数据发送给通道
		channel <- fmt.Sprintf("%s:%v", header, rand.Int31())
		time.Sleep(time.Second)
	}
}

func customer(channel <-chan string) {
	// 不停地获取数据
	for {
		message := <-channel
		fmt.Println(message)
	}
}

func main() {
	// 创建一个字符串类型的通道
	channel := make(chan string)
	// 创建producer()函数的并发goroutine
	go producer("cat", channel)
	go producer("cat", channel)
	go producer("dog", channel)
	// 数据消费函数
	customer(channel)
}
