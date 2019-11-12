package main

import (
	"fmt"
	"time"
)

/**
select语句选择一组可能的send操作和receive操作去处理。它类似switch,但是只是用来处理通讯(communication)操作。
它的case可以是send语句，也可以是receive语句，亦或者default。
receive语句可以将值赋值给一个或者两个变量。它必须是一个receive操作。
最多允许有一个default case,它可以放在case列表的任何位置，尽管我们大部分会将它放在最后。
*/
func fabonacci(c, q chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-q:
			fmt.Println("quit...")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fabonacci(c, quit)

	//超时处理timeout

	chnl := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		chnl <- "result 1"
	}()
	for {
		select {
		case res := <-chnl:
			fmt.Println(res)
			return
		case <-time.After(time.Second * 1):
			fmt.Println("time out 1s")
		}
	}

}
