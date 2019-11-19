package main

import (
	"fmt"
	"time"
)

//并发打印
func printer(c chan int) {
	for {
		data := <-c
		if data == 0 {
			break
		}
		fmt.Println(data)
	}
	//通知main已经结束循环，我搞定了
	c <- 0
}

func main() {
	c := make(chan int, 1000)
	go printer(c)
	for i := 1; i <= 10000; i++ {
		c <- i
		time.Sleep(time.Microsecond)
	}
	//通知并发的printer结束循环，没数据了
	c <- 0
	//等待printer结束,搞定喊我
	<-c
}
