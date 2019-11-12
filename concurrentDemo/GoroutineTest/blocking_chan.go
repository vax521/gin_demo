package main

import (
	"fmt"
	"time"
)

func hello_blocking(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}
func sum(s []int, chnl chan int) {
	sum := 0
	for _, value := range s {
		sum += value
	}
	chnl <- sum
}

func main() {
	done := make(chan bool)
	fmt.Println("start calling")
	go hello_blocking(done)
	<-done
	fmt.Println("main function")

	s := []int{7, 2, 6, -5, 5, 5, 6, 7, 8, 8, 9, 10}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x)
	fmt.Println(y)
}
