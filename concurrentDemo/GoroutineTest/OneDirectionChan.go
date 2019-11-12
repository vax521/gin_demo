package main

import "fmt"

// send only 信道
func sendData(sendCh chan<- int) {
	sendCh <- 10
}

func main() {
	sendCh := make(chan int)
	go sendData(sendCh)
	fmt.Println(<-sendCh)
}
