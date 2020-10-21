package main

import (
	"fmt"
	"time"
)

func timer(duration time.Duration) chan bool {
	ch := make(chan bool)
	go func() {
		time.Sleep(duration)
		ch <- true
	}()
	return ch
}

func main() {
	timer := timer(1 * time.Second)
	for {
		select {
		case <-timer:
			fmt.Println("already 1s")
			return
		}
	}
}
