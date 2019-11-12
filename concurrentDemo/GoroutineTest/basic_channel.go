package main

import (
	"fmt"
	"time"
)

func hello_go(done chan bool) {
	fmt.Println("hello")
	done <- true
}

//关闭信道和使用for range遍历信道
func producer(chnl chan int) {
	for i := 1; i < 10; i++ {
		chnl <- i
		time.Sleep(1 * time.Second)
	}
	close(chnl)
}

func main() {
	done := make(chan bool)
	//协程之间传递信息
	go hello_go(done)
	value := <-done
	fmt.Println(value)
	fmt.Println("main function")

	rece := make(chan int)
	go producer(rece)
	for {
		v, ok := <-rece
		if ok == false {
			break
		}
		fmt.Println("receivied:", v)
	}
}
