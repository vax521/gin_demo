package main

import (
	"fmt"
	"time"
)

func main() {
	/*	timer1 := time.NewTimer(2*time.Second)
		value := <-timer1.C
		fmt.Println(value)
		fmt.Println("Time 1 experied")*/

	timer2 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Time 2 experied")
	}()
	time.Sleep(1 * time.Second)
	stop := timer2.Stop()
	if stop {
		fmt.Println("timer2 stoped")
	}

	ticker1 := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker1.C {
			fmt.Println("Tick at ", t)
		}
	}()
	time.Sleep(10 * time.Second)
}
