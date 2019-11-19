package main

import (
	"fmt"
	"time"
)

func main() {
	//声明一个退出用的通道
	exit := make(chan int)
	fmt.Println("start...")
	//过一秒钟以后调用匿名函数
	time.AfterFunc(time.Second, func() {
		fmt.Println("One Second After")
		//通知main()的goroutine已经结束
		exit <- 0
	})
	//等待结束
	<-exit
}
