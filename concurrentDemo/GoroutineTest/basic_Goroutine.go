package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Printf("Helllo,Goroutine!\n")
}

func main() {
	/*
			启动一个新的协程时，协程的调用会立即返回。与函数不同，程序控制不会去等待 Go 协程执行完毕。在调用 Go 协程之后，程序控制会立即返回到代码的下一行，忽略该协程的任何返回值。
		如果希望运行其他 Go 协程，Go 主协程必须继续运行着。如果 Go 主协程终止，则程序终止，于是其他 Go 协程也不会继续运行。
	*/
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}
